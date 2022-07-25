# gotemplate - шаблон репозитория GoForOps

структура папок для выполнения домашних заданий и сдачи их на проверку

## Как работает grader

Grader-ом у нас будет gitlab-ci и его раннеры. В файлы, в указанных ниже папках, вносить изменения не следует.

```bash
├── .gitlab-ci      # директория с файлами gitlab-ci: общие переменные, шаги тестирования
│   ├── default.yml # Общие переменные
│   ├── prepare.yml # Helper для вывода переменных окружения gitlab-a
│   └── test.yml    # Шаги тестирования с правилами запуска
├── .gitlab-ci.yml  # основной gitlab-ci файл
```

Правила содержат основое условие - изменения в файлах определенной директории

```yaml
rules:
    - changes:  # Include the job and set to when:manual if any of the follow paths match a modified file.
        - Exercises/2.Basics/1/*.go
```

Основной скрипт - элементарная проверка форматирования и запуск тестов

```yaml
  script:
    - if [[ -n $(gofmt -l Exercises/2.Basics/1/) ]]; then echo "Необходимо отформатировать код при помощи gofmt" && exit 1; fi
    - go test -v Exercises/2.Basics/1/*.go
    - echo "Домашка 2.1 пройдена! Можно сдавать на ревью!"
```

Вы можете запустить тесты у себя локально, подсмотрев команду в `.gitlab-ci/test.yml`, затем пушить уже на сервер.

## Как использовать данный репозиторий

Форкнуть репозиторий себе, namespace выбираем go_s<ваш номер студента>  
Отвязать форк от родителя: `Settings/General/Advanced/Expand` -> `Remove fork relationship`

Подробнее описано в шаге 1.4.3

- Для заданий на закрепления навыков код писать в соответствующем разделе в папке `Exercises`.  
- Для заданий касающихся разработки микросервиса код писать в папке `Service`.  

2 папки являются шаблонами для копирования из `Exercises` в `Service` на соответствующем этапе: `Exercises/8.Kubernetes/2` и `Exercises/8.Kubernetes/3`
