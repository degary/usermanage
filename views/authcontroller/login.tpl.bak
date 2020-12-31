<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Login</title>
    <link rel="stylesheet" type="text/css" href="/static/bootstrap-4.5.3-dist/css/bootstrap.css">
    <link rel="stylesheet" type="text/css" href="/static/css/login.css">
</head>
<body>
    <div class="container">
        <div class="row">
            <div class="col-4 offset-4">
                <div class="login-form shadow-lg p-3 mb-5 bg-white rounded">
                    <h2 class="text-center">用户登录</h2>
                    <hr />
                    <form action="/auth/login/" method="post">
                        <div class="form-group">
                            <label class="sr-only">用户名: </label>
                            <input class="form-control" name="username" type="text" value="" placeholder="用户名">
                        </div>
                        <div class="form-group">
                            <label class="sr-only">密码: </label>
                            <input class="form-control" name="password" type="password" value="" placeholder="密码">
                        </div>
                        {{ if .errors.default }}
                        <p class="text-danger">{{ index .errors.default 0 }}</p>
                        {{ end }}
                        <div class="form-group">
                            <input class="btn btn-primary btn-block" type="submit" value="登录">
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</body>
</html>