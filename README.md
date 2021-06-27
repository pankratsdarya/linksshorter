# Курсовой проект по курсу Backend-разработка на Go. Уровень 1.

# Назначение:
URL shortener (“сокращатель ссылок”) - это сервис, позволяющий пользователю генерировать  для произвольного URL’a его короткую версию, которую удобно вставлять в различные публикации, сообщения, новости, промо-материалы и так далее. Также сервис позволяет получать статистику переходов по каждому сгенерированному URL’у, что будет полезно, если его владелец захочет узнать сколько людей перешло по короткой ссылке из конкретного источника.

# Принцип работы:
Приложение условно можно разделить на 3 части:
1. Frontend - пользовательский интерфейс. Должен содержать:
a. окно для вставки длинных ссылок;
б. кнопку "Сгенерировать" для генерации коротких ссылок;
в. окно для коротких ссылок - полученных после генерации, прописанных при переходе с другого адреса или вставленных из буфера;
г. кнопку "Перейти" для перехода по короткой ссылке;
д. окно со статистикой.
Реализация не требуется.
2. Хранилище - на начальном этапе - файл. Должно содержать статистику:
а. пары длинная-короткая ссылки;
б. количество переходов по каждой паре;
в. время последнего обращения к ссылке (удалять ссылку из хранилища, если к ней не обращались год);
Реализация пока малопонятна. _to be developed_
3. Backend - вся математика приложения. Должен содержать:
а. интерфейс с хранилищем;
б. интерфейс с frontend;
в. пакет с генерацией ссылок; 
г. пакет обработки переходов по ссылкам, сбор статистики;
д. пакет работы с хранилищем;
е. пакет поддержания жизни хранилища (пока просто чистит ссылки, после перехода на бд - проверяет соединение с ней);
ж. логирование;
Реализация пока малопонятна. _to be developed_

# Примечание к текущей версии
Текущая версия программы работает с файлом в качестве хранилища. Для корректной работы 
можно запускать только одну версию приложения. Данное ограничение будет устранено после
перехода на базу данных в качестве хранилища.