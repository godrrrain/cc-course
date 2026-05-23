# Архитектура Dart Компилятора

## Общая архитектура

```
Исходный код (hello.dart)
         ↓
    ┌────────────┐
    │   Лексер   │ (Dart2Lexer.g4)
    │ (ANTLR4)   │
    └────────────┘
         ↓
     Токены
         ↓
    ┌────────────┐
    │   Парсер   │ (Dart2Parser.g4)
    │ (ANTLR4)   │
    └────────────┘
         ↓
    Abstract Syntax Tree (AST)
         ↓
    ┌────────────┐
    │  Visitor   │ (internal/visitor/)
    │ Паттерн    │
    └────────────┘
         ↓
  LLVM IR (.ll файл)
         ↓
    ┌────────────┐
    │   Clang    │
    │            │
    └────────────┘
         ↓
  Исполняемый файл (.exe / бинарь)
```

## Процесс компиляции

### Этап 1: Лексический анализ

**Входные данные:** Исходный Dart код

**Процесс:**
1. Чтение исходного файла символ за символом
2. Группировка символов в лексемы (токены)
3. Распознавание ключевых слов, операторов, идентификаторов, литералов

**Выходные данные:** Поток токенов

**Пример:**
```
Входной код:  void main() { int a = 5; }
Токены:       VOID IDENTIFIER LPAREN RPAREN LBRACE INT IDENTIFIER ASSIGN NUMBER RBRACE
```

### Этап 2: Синтаксический анализ

**Входные данные:** Поток токенов

**Процесс:**
1. Чтение токенов в соответствии с грамматикой Dart
2. Построение Abstract Syntax Tree (AST)
3. Проверка грамматических правил

**Выходные данные:** AST

**Пример AST для `int a = 5;`:**
```
VariableDeclaration
├── type: "int"
├── name: "a"
└── initializer: NumericLiteral(5)
```

### Этап 3: Семантический анализ и кодогенерация

**Входные данные:** AST

**Процесс:**
1. Обход AST с помощью visitor паттерна
2. Управление областями видимости переменных
3. Генерация LLVM IR инструкций
4. Создание функций, блоков, инструкций

**Выходные данные:** LLVM IR

**Пример LLVM IR для `int a = 5;`:**
```llvm
%1 = alloca i64
store i64 5, i64* %1
```

### Этап 4: Компиляция LLVM

**Входные данные:** LLVM IR файл (.ll)

**Процесс:**
1. LLVM оптимизирует IR
2. Генерирует машинный код для целевой платформы
3. Объединяет с runtime библиотеками
4. Создает исполняемый файл

**Выходные данные:** Исполняемый файл

## Структура кода

### Основной поток выполнения

```
cmd/compiler/main.go
    ↓
internal/compiler/compiler.go::Compiler()
    ↓
    ├─→ antlr.NewFileStream()      // Чтение файла
    ├─→ parser.NewDart2Lexer()     // Создание лексера
    ├─→ parser.NewDart2Parser()    // Создание парсера
    ├─→ parser.CompilationUnit()   // Парсирование
    ├─→ visitor.NewIRVisitor()     // Создание visitor'а
    ├─→ visitor.Visit(tree)        // Обход AST
    ├─→ visitor.WriteToFile()      // Запись LLVM IR
    └─→ exec.Command("clang")      // Компиляция
```

### Модули

#### 1. cmd/compiler/main.go

Точка входа программы. Обрабатывает:
- Параметры командной строки (-i, -o, -d)
- Вызов основного компилятора
- Обработка ошибок

```go
func main() {
    flag.Parse()
    err := compiler.Compiler(*input, *output, *isDOT)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
```

#### 2. internal/compiler/compiler.go

Основная логика компиляции:

```go
func Compiler(input, output string, isDOT bool) error {
    fileStream, _ := antlr.NewFileStream(input)
    
    lexer := parser.NewDart2Lexer(fileStream)
    tokens := antlr.NewCommonTokenStream(lexer, ...)
    p := parser.NewDart2Parser(tokens)
    
    tree := p.CompilationUnit()
    
    v := visitor.NewIRVisitor()
    v.Visit(tree)
    
    v.WriteToFile(output + ".ll")
    
    exec.Command("clang", output+".ll", "-o", output+".exe").Run()
}
```

#### 3. internal/visitor/visitor.go

Главный visitor класс, наследует от `BaseDart2ParserVisitor`

**Основные методы:**
- `VisitCompilationUnit()` - обработка верхнеуровневых элементов
- `VisitTopLevelDeclaration()` - обработка деклараций
- `VisitFunctionDeclaration()` - обработка функций
- `VisitStatement()` - обработка операторов

**Ключевые поля:**
```go
type IRVisitor struct {
    Module      *ir.Module    // LLVM модуль
    currentFunc *ir.Func      // Текущая функция
    currentScope *Scope       // Область видимости
    currentBlock *ir.Block    // Текущий блок кода
    Errors      []error       // Ошибки
}
```

#### 4. internal/visitor/statement.go

Обработка операторов:
- `VisitIfStatement()` - условные операторы
- `VisitWhileStatement()` - циклы while
- `VisitForStatement()` - циклы for
- `VisitReturnStatement()` - return операторы
- `VisitBreakStatement()` - break операторы

#### 5. internal/visitor/expression.go

Обработка выражений:
- `VisitPrimary()` - первичные выражения (переменные, литералы)
- `VisitBinaryExpression()` - бинарные операторы (+, -, *, /)
- `VisitLiteral()` - литеральные значения
- `VisitCallExpression()` - вызовы функций

#### 6. internal/visitor/models.go

Данные модели:
- `Scope` - область видимости переменных
- `VariableInfo` - информация о переменной
- `FuncWithParams` - информация о функции с параметрами

```go
type VariableInfo struct {
    Name      string           // Имя переменной
    LLVMValue value.Value      // LLVM значение
    Type      types.Type       // LLVM тип
    Direction ParamDirection   // Направление (in/out/inout)
    Bounds    []ArrayBounds    // Границы для массивов
}

type Scope struct {
    parent *Scope
    vars   map[string]*VariableInfo
    meta   map[string]interface{}
}
```

## Управление областью видимости

Visitor использует иерархию scopes для управления видимостью переменных:

```
Global Scope (functions, global variables)
    ↓
Function Scope
    ├─→ Block Scope 1 (if statement)
    │   └─→ Variable: x (local)
    └─→ Block Scope 2 (else statement)
        └─→ Variable: y (local)
```

**Методы управления:**
- `enterScope()` - создать новую область видимости
- `exitScope()` - выйти из текущей области видимости
- `currentScope.Get(name)` - получить переменную
- `currentScope.Set(name, info)` - установить переменную

## LLVM IR генерация

### Типы данных

Преобразование типов Dart в LLVM:

```
Dart Type    →    LLVM Type
int          →    i64
double       →    double
bool         →    i1
String       →    i8* (указатель на строку)
List<T>      →    массив T
void         →    void
```

### Инструкции

Примеры генерирования LLVM инструкций:

```go
// Присваивание
v.currentBlock.NewStore(rhsVal, lhsVal)

// Сложение
v.currentBlock.NewAdd(lhs, rhs)

// Условный переход
v.currentBlock.NewCondBr(cond, thenBlock, elseBlock)

// Безусловный переход
v.currentBlock.NewBr(nextBlock)

// Возврат из функции
v.currentBlock.NewRet(returnVal)
```

## Пример: Компиляция простой программы

**Входной Dart код:**
```dart
int add(int a, int b) {
  return a + b;
}

void main() {
  int result = add(3, 5);
  print(result);
}
```

**Сгенерированный LLVM IR:**
```llvm
declare i32 @printf(i8*, ...)

define i64 @add(i64 %a, i64 %b) {
entry:
  %add = add i64 %a, %b
  ret i64 %add
}

define void @main() {
entry:
  %call = call i64 @add(i64 3, i64 5)
  %call1 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @.str, i32 0, i32 0), i64 %call)
  ret void
}

@.str = private unnamed_addr constant [3 x i8] c"%d\00", align 1
```

## Обработка ошибок

Ошибки собираются в `IRVisitor.Errors`:

```go
if err != nil {
    v.Errors = append(v.Errors, fmt.Errorf("..."))
}

// После обхода дерева
if len(v.Errors) > 0 {
    return fmt.Errorf("compilation errors: %v", v.Errors)
}
```

## Оптимизация

В текущей версии оптимизация выполняется на уровне LLVM:

```bash
opt -O2 hello.ll -o hello.opt.ll
clang hello.opt.ll -o hello
```

## Производительность

- **Лексический анализ**: O(n) где n - длина кода
- **Синтаксический анализ**: O(n) для большинства языков
- **Семантический анализ**: O(n) для однопроходного visitor
- **Кодогенерация**: O(n)

Общая сложность: **O(n)**

## Расширяемость

Компилятор легко расширяется:

1. **Добавить ключевое слово/оператор:**
   - Добавить в грамматику (Dart2Lexer.g4, Dart2Parser.g4)
   - Сгенерировать парсер: `make antlr4`
   - Реализовать visitor метод

2. **Добавить новый тип:**
   - Определить в грамматике
   - Добавить в VariableInfo
   - Реализовать преобразование в LLVM тип

3. **Добавить оптимизацию:**
   - Реализовать проход через AST
   - Трансформировать узлы
   - Встроить в visitor

## Ссылки

- [ANTLR4 dokumentace](https://www.antlr.org/wiki/display/ANTLR4/Documentation)
- [LLVM Language Reference Manual](https://llvm.org/docs/LangRef/)
- [Go LLVM Bindings](https://github.com/llir/llvm)
- [Dart Language Specification](https://dart.dev/guides/language/specifications)
