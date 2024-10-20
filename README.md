# Go-Watcher: Защитник вашего веб-сервера

**Go-Watcher** - это легковесный инструмент безопасности, написанный на Go, который отслеживает подозрительную активность на вашем веб-сервере и блокирует ее. Он анализирует логи доступа Nginx в режиме реального времени, выявляя и блокируя потенциально вредоносные запросы.

## Ключевые возможности

- **Мониторинг логов Nginx в реальном времени**: Go-Watcher постоянно отслеживает логи доступа Nginx на наличие подозрительных запросов.
- **Обнаружение атак**: Используя настраиваемые правила и регулярные выражения, Go-Watcher может идентифицировать различные типы атак, такие как:
  - Сканирование портов
  - SQL-инъекции
  - Межсайтовый скриптинг (XSS)
  - Брутфорс
- **Блокировка IP-адресов**: При обнаружении подозрительной активности Go-Watcher может автоматически блокировать IP-адрес злоумышленника, добавляя его в черный список Nginx.
- **Гибкая конфигурация**: Вы можете настроить Go-Watcher для:
  - Мониторинга определенных путей
  - Игнорирования определенных IP-адресов
  - Настройки чувствительности обнаружения
- **Ведение журнала событий**: Go-Watcher ведет подробный журнал всех обнаруженных событий, включая:
  - Дату и время
  - IP-адрес злоумышленника
  - Заблокированный URL-адрес
  - Тип атаки
- **Простота использования**: Go-Watcher прост в настройке и использовании, что делает его доступным даже для начинающих пользователей.

## Установка

### Шаг 1: Скачайте последнюю версию Go-Watcher

Вы можете найти последнюю версию на [странице релизов проекта](ссылка_на_релизы).

### Шаг 2: Сконфигурируйте Go-Watcher

1. Отредактируйте файл конфигурации `config.json`:
   - Задайте правила обнаружения.
   - Настройте другие параметры.

   Пример конфигурации `config.json`:

### Шаг 3: Запустите Go-Watcher

Запустите исполняемый файл Go-Watcher:

```bash
./go-watcher
```

## Пример использования

Предположим, вы хотите защитить свой веб-сервер от атак брутфорс на страницу авторизации. Вы можете настроить Go-Watcher для мониторинга запросов к `/wp-login.php` и блокировки IP-адресов, которые делают более 5 неудачных попыток входа в систему в течение 1 минуты. Просто добавьте соответствующие правила в ваш файл `config.json`, как показано в примере выше.
