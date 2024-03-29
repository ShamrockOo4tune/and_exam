## For english scroll down
# Экзаменационное задание курсов DevOps компании Andersen  

Цель проекта: наладить и автоматизировать CI/CD приложения используя контейнеризацию docker.
Приложение: 2 веб приложения на различных языках.
Я выбрал Python - Flask (интерпретируемый) и Golang - Gin (компилируется). Файлы приложений хранятся в отдельных папках: py_program_1 and go_program_2 

Структура репозитория проекта:  
.  
├── .github/workflows  
│   └── docker-image.yml ## имя файла не отражает сути и намерений его использования  
├── dockerfiles  
│   ├── GoDockerfile  
│   └── PyDockerfile  
├── go_program_2  
│   ├── go.mod  
│   ├── go.sum  
│   └── main.go  
├── py_program_1  
│   ├── app.py  
│   └── requirements.txt  
├── project_slides.pdf  
└── README.md  

В качестве CI/CD используется GitHub actions. GitHub actions инициируется файлом .yml в директории .github/workflows. Он настраивает автоматический запуск процесса при коммите на ветку master.

Реализованная схема CI/CD схематично изображена на втором слайде project_slides.pdf  
Рабочий процесс, описанный в .yml файле требует наличия предустановленных секретов GitHub и переменных окружения:  
  Учетные данные AWS CLI:  
    AWS_ACCESS_KEY_ID  
    AWS_SECRET_ACCESS_KEY  
  Учетные данные Docker Hub Registry:  
    DOCKER_HUB_ACCOUNT  
    DOCKER_HUB_PASSWORD  
  Учетные данные Telegram бота:  
    TELEGRAM_TO    
    TELEGRAM_TOKEN  
  Данные AWS EC2 инстансов:  
    AWS_REGION_NAME  
    and_exam_webserver1_id  (id интанса EC2 сохраняется постоянным на стороне AWS)   
    and_exam_webserver2_id   
  Составные имена Docker образов: 
    py_docker-image        (Составное имя образа докер контейнера, уникальное для каждого запуска процесса при коммите)
    go_docker-image
<br>
Как только рабочий процесс инициирован триггером, начинается сборка докер образов на основе PyDockerfile и GoDockerfile. Сборка происходит на инфраструктуре предоставляемой Github.

Готовые образы контейнеров публикуются в Dockerhub registry. Учетные данные хранятся в скрытом виде как Github secrets.

В Docker Hub registry установлена политика удаления редкоиспользуемых образов. Установка пользовательских политик не реализована.

Следующий этап рабочего процесса: проверка что контейнеры (и приложения в них) работают как задуманно. Образы копируются из dockerhub и запускаются на инфраструктуре предоставляемой Github. Утилита CURL используется для получения ответа от веб приложений и позволяет убедиться в доступности сгенерированных страниц.

Следующий этап - разворачивание приложений на продакшн. Происходит запуск предварительно сконфигурированных инстансов AWS EC2, доставка образов и запуск контейнеров.

Уведомления по рабочему процессу аккумулируются в логах, которые доступны к просмотру и копированию из вебинтерфейса Github.
Некоторые уведомления направляются в Телеграм бот для получения оперативной сводки о ходе выполнения процесса.





# Andersen DevOps courses exam project
<br>  
The project is to setup and automate application CI/CD using docker containers  
Application: 2 webapp on two different languages.  
I have selected Python - Flask (interpreted) and Golang - Gin (being compiled). The apps files are staged in separate folders: py_program_1 and   go_program_2   
<br>  
The structure of the project repo: <br>  
. <br>   
├── .github/workflows <br>    
│   └── docker-image.yml ## file name doesnt reflect actual intentions  <br>  
├── dockerfiles    <br>
│   ├── GoDockerfile    <br>
│   └── PyDockerfile    <br>
├── go_program_2    <br>
│   ├── go.mod    <br>
│   ├── go.sum    <br>
│   └── main.go    <br>
├── py_program_1    <br>
│   ├── app.py    <br>
│   └── requirements.txt    <br>
├── project_slides.pdf    <br>
└── README.md    <br>
<br>  
The CI/CD tool used is GitHub actions. Github actions are initiated by the .yml file in the .github/workflows directory. It is set to run the workflow on push to master branch of the repo.  
<br>  
Implemented CI/CD scheme is presented on slide #2 of project_slides.pdf    
The workflow described in .yml file requires the repo to have pre set secrets and environment variables:  
  AWS CLI access credentials:  
    - AWS_ACCESS_KEY_ID  
    - AWS_SECRET_ACCESS_KEY  
  Docker Hub Registry credentials:  
    - DOCKER_HUB_ACCOUNT  
    - DOCKER_HUB_PASSWORD  
  Telegram bot credentials:  
    - TELEGRAM_TO            
    - TELEGRAM_TOKEN   
  AWS EC2 instances related data:   
    - AWS_REGION_NAME  
    - and_exam_webserver1_id (EC2 instance ID kept constant on AWS)   
    - and_exam_webserver2_id  
  Docker images composite names:  
    - py_docker-image    (Composite docker image name, unique per every commit)  
    - go_docker-image  
<br>  
Once the workflow triggered it starts to build docker images based on PyDockerfile and GoDockerfile. Image assembly is done on the Github provided infrastructure.  
<br>  
Complete images are being pushed to Dockerhub registry. Credentials are kept secret as Github repo secrets   <br>
<br>   
Next workflow job is to check if the containers and the apps in this containers are working properly. The images are pulled from dockerhub and run in github provided infrastructure. CURL utility is used to recieve the webapps response and ensure the apps are generating accessible web pages.  
<br>   
Next workflow job is to deploy on production. Invoke preconfigured AWS infrastructure (EC2 instances), deliver images and start containers. Once ready the accessible IP:port info is provided and web pages can be visited.   
<br>   
The workflow notifications are accumulated in logfiles which are accessible from github actions web interface.  
Some notifications are also redirected to Telegram bot to recieve operative updates on the worflow progress.  
