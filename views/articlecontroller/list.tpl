<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>文章列表</title>
    <link rel="stylesheet" type="text/css" href="/static/bootstrap-4.5.3-dist/css/bootstrap.css">
    <link rel="stylesheet" type="text/css" href="/static/css/layout.css">
    <!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
    <script src="/static/bootstrap-4.5.3-dist/js/jquery-3.5.1.min.js"></script>

    <!-- bootstrap.bundle.min.js 用于弹窗、提示、下拉菜单，包含了 popper.min.js -->
    <script src="/static/bootstrap-4.5.3-dist/js/bootstrap.bundle.min.js"></script>

    <!-- 最新的 Bootstrap4 核心 JavaScript 文件 -->
    <script src="/static/bootstrap-4.5.3-dist/js/bootstrap.min.js"></script>
</head>
<body>

    {{ template "basehtml/nav.html" . }}
    <div class="container-fluid">
        <div class="row">
            <div class="col-12 table-btn">
                <a class="btn btn-primary" href="/article/add/">新建</a>
            </div>
        </div>
        <table class="table table-striped">
            <thead>
            <tr>
                <th scope="col">ID</th>
                <th scope="col">书名</th>
                <th scope="col">作者</th>
                <th scope="col">价钱</th>
                <th scope="col">操作</th>


            </tr>
            </thead>
            <tbody>
            {{ range .articles }}
                <tr>
                    <td>{{ .Id }}</td>
                    <td>{{ .Name }}</td>
                    <td>{{ .Author }}</td>
                    <td>{{ .Price }}</td>
                    <td>
                        <a class="btn btn-sm btn-warning" href="/article/update/?id={{ .Id }}">编辑</a>
                        <a class="btn btn-sm btn-danger" href="/article/delete/?id={{ .Id }}">删除</a>
                    </td>
                </tr>
            {{ end }}
            </tbody>
        </table>
        {{ template "basehtml/pagetion.html" }}
    </div>



</body>
</html>