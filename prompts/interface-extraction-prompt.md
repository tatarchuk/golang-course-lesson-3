# Промпт для виокремлення інтерфейсу (Розділ 2)

Вставте нижче точний текст промпту, який ви розробили для Розділу 2
домашньої роботи. Промпт має явно містити:

- **роль** ШІ (наприклад, "Ти — досвідчений Go-розробник, що
  спеціалізується на чистій архітектурі...");
- **контекст** — наведений у завданні код `OrderService`;
- **завдання** — виокремити мінімальний інтерфейс `OrderStore`,
  переписати `OrderService` на dependency injection, згенерувати мок;
- **обмеження** — маленький інтерфейс (приказка Роба Пайка),
  конкретний стиль коду, тощо;
- **бажаний формат виводу** (наприклад, "поверни лише Go-код у трьох
  блоках: інтерфейс, сервіс, мок").

## Мій промпт
You are the Senior Go developer, and know how to make a clean code, we need to refactor order.go file. 
Please suggest me refactored file content as output taking into account two points: interface OrderStore now using a      
single method to execute DB queries, please split it into atomic CRUD operations; 
OrderService should acquire dependency through constructor using dependency injection 
