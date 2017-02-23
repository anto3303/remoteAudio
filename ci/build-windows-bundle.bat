mkdir %GOPATH%\src\github.com\dh1tw\remoteAudio\release\
%MSYS_PATH%\usr\bin\bash -lc "cp /mingw%MSYS2_BITS%/**/*.dll /c/gopath/src/github.com/dh1tw/remoteAudio/release/"
%MSYS_PATH%\usr\bin\bash -lc "pacman --noconfirm --needed -Sy sed" > nul
%MSYS_PATH%\usr\bin\bash -lc "cd /c/gopath/src/github.com/twstrike/coyim && ci/release"
%MSYS_PATH%\usr\bin\bash -lc "cd /c/gopath/src/github.com/dh1tw/remoteAudio/release && 7z a -tzip remoteAudio.zip *"
%MSYS_PATH%\usr\bin\bash -lc "cd /c/gopath/src/github.com/dh1tw/remoteAudio/release && rm *.dll"
%MSYS_PATH%\usr\bin\bash -lc "cd /c/gopath/src/github.com/dh1tw/remoteAudio/release && rm *.exe"
xcopy %GOPATH%\src\github.com\dh1tw\remoteAudio\release\remoteAudio.zip %APPVEYOR_BUILD_FOLDER%\ /e /i > nul