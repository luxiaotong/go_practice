FROM python:3.8-alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update -no-cache && \
    apk add -no-cache python3-tkinter ghostscript make automake gcc g++ \
    linux-headers musl libxml2-dev libxslt-dev libffi-dev \
    musl-dev libgcc openssl-dev jpeg-dev zlib-dev freetype-dev build-base \
    lcms2-dev openjpeg-dev cmake clang clang-dev ninja \
    zlib jpeg libjpeg freetype openjpeg curl

RUN ln -s /usr/include/locale.h /usr/include/xlocale.h

RUN curl -L https://github.com/opencv/opencv/archive/4.2.0.tar.gz | tar zx && \
    cd opencv-4.2.0 && mkdir build && cd build && \
    cmake -G Ninja \
    -D CMAKE_BUILD_TYPE=RELEASE \
    -D CMAKE_INSTALL_PREFIX=/usr/local \
    -D WITH_FFMPEG=NO \
    -D WITH_IPP=NO \
    -D PYTHON_EXECUTABLE=/usr/local/bin/python \
    -D WITH_OPENEXR=NO ..
RUN ninja && ninja install

# RUN ln -s /usr/local/lib/python3.8/site-packages/cv2/python-3.8/cv2.cpython-38-x86_64-linux-gnu.so /usr/lib/python3.8/site-packages/cv2.so
RUN ln -s /usr/lib/libgs.so.9 /usr/lib/libgs.so


RUN pip install camelot-py[cv] -i https://pypi.tuna.tsinghua.edu.cn/simple