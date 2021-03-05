<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">

    <title>Gin Hello</title>
    <link rel="icon" href="assets/img/favicon.ico">
    <link rel="stylesheet" href="assets/css/bootstrap.min.css">
    <link rel="stylesheet" href="assets/css/style.css">
    
    <script rel="script" src="/assets/js/jquery-3.5.1.min.js"></script>
    <script rel="script" src="/assets/js/bootstrap.min.js"></script>
</head>
<body>
    <header>{{ template "navi" }}</header>

    {{ $type := .content }}
    {{ if eq $type "blank" }}
        {{ template "blank" . }}
    {{ else if eq $type "login" }}
        {{ template "login" . }}
    {{end}}
</body>
</html>