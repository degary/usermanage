<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>文章列表</title>
    <link rel="stylesheet" type="text/css" href="/static/bootstrap-4.5.3-dist/css/bootstrap.css">
    <link rel="stylesheet" type="text/css" href="/static/css/layout.css">
</head>
<body>

    {{ template "basehtml/nav.html" . }}
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
        <tr>
            <th scope="row">1</th>
            <td>Mark</td>
            <td>Otto</td>
            <td>Otto</td>
            <td>Otto</td>
        </tr>
        </tbody>
    </table>
    {{ template "basehtml/pagetion.html" }}


</body>
</html>