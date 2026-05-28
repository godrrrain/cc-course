void main() {
  List<int> numbers = [1, 2, 3, 4, 5];
  int n = 5;
  int temp = 0;
  
  print("Array:");
  for (int i = 0; i < n; i++) {
    print(numbers[i]);
  }
  
  for (int i = 0; i < n / 2; i++) {
    temp = numbers[i];
    numbers[i] = numbers[n - 1 - i];
    numbers[n - 1 - i] = temp;
  }
  
  print("Reversed Array:");
  for (int i = 0; i < n; i++) {
    print(numbers[i]);
  }
}