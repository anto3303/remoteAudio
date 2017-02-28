mkdir c:\release\
%MSYS_PATH%\usr\bin\bash -lc "cp /mingw%MSYS2_BITS%/**/libogg-0.dll /c/release/"
%MSYS_PATH%\usr\bin\bash -lc "cp /mingw%MSYS2_BITS%/**/libopus-0.dll /c/release/"
%MSYS_PATH%\usr\bin\bash -lc "cp /mingw%MSYS2_BITS%/**/libopusfile-0.dll /c/release/"
%MSYS_PATH%\usr\bin\bash -lc "cp /mingw%MSYS2_BITS%/**/libportaudio-2.dll /c/release/"
%MSYS_PATH%\usr\bin\bash -lc "cp /mingw%MSYS2_BITS%/**/libsamplerate-0.dll /c/release/"
REM %MSYS_PATH%\usr\bin\bash -lc "pacman --noconfirm --needed -Sy sed" > nul
REM %MSYS_PATH%\usr\bin\bash -lc "cd /c/gopath/src/github.com/dh1tw/remoteAudio && ci/release"
REM %MSYS_PATH%\usr\bin\bash -lc "cp /c/gopath/src/github.com/dh1tw/remoteAudio/remoteAudio.exe /c/release/"
REM %MSYS_PATH%\usr\bin\bash -lc "cd /c/gopath/src/github.com/dh1tw/remoteAudio/release && 7z a -tzip remoteAudio.zip *"
REM %MSYS_PATH%\usr\bin\bash -lc "cd /c/gopath/src/github.com/dh1tw/remoteAudio/release && rm *.dll"
REM %MSYS_PATH%\usr\bin\bash -lc "cd /c/gopath/src/github.com/dh1tw/remoteAudio/release && rm *.exe"
REM xcopy %GOPATH%\src\github.com\dh1tw\remoteAudio\release\remoteAudio.zip %APPVEYOR_BUILD_FOLDER%\ /e /i > nul