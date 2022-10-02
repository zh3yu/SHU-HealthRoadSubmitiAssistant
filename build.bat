title "Filing Assistant WINDOS compilation script "
echo "Ensure that the GO runtime environment is installed"
if exist releases (
rd /s/q .\releases
) 
md releases\SHU-HealthRoadSubmitiAssistant\Configuration
md releases\SHU-HealthRoadSubmitiAssistant\Resources
xcopy \Configuration releases\SHU-HealthRoadSubmitiAssistant\Configuration
xcopy \Resources releases\SHU-HealthRoadSubmitiAssistant\Resources
go build -o releases\SHU-HealthRoadSubmitiAssistant\SHUHR.exe  
cd releases
tar -cvf SHU-HealthRoadSubmitiAssistant-releases.zip SHU-HealthRoadSubmitiAssistant  
echo DONE
cd ../
pause