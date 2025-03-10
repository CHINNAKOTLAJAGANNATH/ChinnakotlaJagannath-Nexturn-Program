pipeline {
    agent any

    environment {
        FLASK_PORT = "8000"
    }

    stages {
        stage('Clone Repository') {
            steps {
                echo 'Cloning the React ToDo Application repository...'
                git branch: 'main', url: 'https://github.com/feelinganubhav/Library-Management-System.git'
            }
        }
        
        stage('Setup Virtual Environment') {
            steps {
                echo "Setting up the environment..."

                // Create virtual environment if it doesn't exist
                bat """
                if not exist venv (
                    python -m venv venv
                )
                """
                bat "venv\\Scripts\\activate && python -m pip install --upgrade pip"
            }
        }

        stage('Install Dependencies') {
            steps {
                echo 'Installing required Python packages...'
                bat """
                call venv\\Scripts\\activate
                if exist requirements.txt (
                    pip install -r requirements.txt
                ) else (
                    echo "ERROR: requirements.txt not found!"
                    exit 1
                )
                call venv\\Scripts\\deactivate
                """
            }
        }

        
        stage('Run All Unit Tests') {
            steps {
                echo 'Running unit tests with pytest...'
                bat """
                call venv\\Scripts\\activate
                set PYTHONPATH=%CD%
                pytest tests/app_pytest.py
                call venv\\Scripts\\deactivate
                """
            }
        }

        stage('Start Flask Server') {
            steps {
               echo 'Starting Flask app using Waitress (Gunicorn alternative for Windows)...'
                bat "start /B waitress-serve --listen=127.0.0.1:%FLASK_PORT% app:app"
                sleep 5
                echo 'Waiting for a few seconds to allow the Flask app to start'
                bat "powershell -Command 'Start-Sleep -Seconds 3'"

            }
        }

        stage('Post-Deployment Testing') {
            steps {
                echo 'Verifying Flask App Routes...'

                bat """
                curl -X GET http://127.0.0.1:%FLASK_PORT%/ -s -o response_home.txt -w "HTTP Code: %%{http_code}\\n" | findstr "200"
                if %ERRORLEVEL% NEQ 0 (
                    echo ERROR: Home route verification failed!
                    taskkill /F /IM python.exe
                    exit 1
                )
                type response_home.txt
                echo Home route verified successfully!
                
                curl -X GET -H "Authorization: you-will-never-guess" http://127.0.0.1:%FLASK_PORT%/books/ -s -o response_books.txt -w "HTTP Code: %%{http_code}\\n" | findstr "200" 
                if %ERRORLEVEL% NEQ 0 (
                    echo ERROR: Books route verification failed!
                    taskkill /F /IM python.exe
                    exit 1
                )
                type response_books.txt
                echo Books route verified successfully!

                curl -X GET -H "Authorization: you-will-never-guess" http://127.0.0.1:%FLASK_PORT%/members/ -s -o response_members.txt -w "HTTP Code: %%{http_code}\\n" | findstr "200" 
                if %ERRORLEVEL% NEQ 0 (
                    echo ERROR: Members route verification failed!
                    taskkill /F /IM python.exe
                    exit 1
                )
                type response_members.txt
                echo Members route verified successfully!
                """
            }
        }
    }

    post {
        always {
            echo 'Cleaning up any Running Server...'
            bat "taskkill /F /IM python.exe /T || exit 0"
        }
        success {
            echo 'Flask App Deployment Successful! 🎉'
        }
        failure {
            echo 'Pipeline Failed! Check the logs for errors. ❌'
        }
    }
}