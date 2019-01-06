
node ("master") {
    stage("checkout") {
        checkout scm
    }

    stage("get") {
        withEnv(["PATH=${env.PATH}:/usr/lib/go-1.10/bin",
                 "GOPATH=${env.HOME}/go",
                 "GOBIN=${env.HOME}/go/bin",
        ]){
            sh "go get"
        }
    }

    stage("build") {
        withEnv([
                "PATH=${env.PATH}:/usr/lib/go-1.10/bin",
                "GOPATH=${env.HOME}/go",
                "GOBIN=${env.HOME}/go/bin",
        ]){
            sh "go build -o ${env.HOME}/.terraform.d/plugins/terraform-provider-appstream_v0.0.1"
        }
    }
}