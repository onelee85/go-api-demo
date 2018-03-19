@echo off

echo ================git pull=======================
@call git pull

sleep 1

@call bee pack -be GOOS=linux

if %ERRORLEVEL% EQU 0 (
        echo ================SUCCESS=======================
        call:pubApi 39.106.130.135 "publish to 39.106.130.135"
        echo ================SUCCESS=======================
        sleep 15
) else (
        COLOR C
	    echo -------         !! FAILD !!      -------------
        pause
)

exit

:pubApi
scp user.tar.gz webserver@%~1:~/go-user
ssh webserver@%~1  "source /etc/profile;cd ~/go-user; tar -xvf user.tar.gz;./restart.sh"