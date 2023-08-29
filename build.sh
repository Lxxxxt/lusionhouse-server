
GOOS=linux GOARCH=amd64 go build -o main .

id=`uuid`
img=lxxxxd/lusionhouse:$id

function runbuild() {
    docker build -t $img . &&\
    docker push $img #&& \
}

function clearimg() {
    docker rmi $img
}

runbuild

if [ $? == 0 ]; then

    clearimg
    

    echo image is : $img
    echo ''

    echo '<<<<<<< FINISHED 🎉🎉🎉 >>>>>'

    echo ''
else

    echo ''

    echo '<<<<<<< FAILED ❌❌❌ >>>>>'

    echo ''

fi

