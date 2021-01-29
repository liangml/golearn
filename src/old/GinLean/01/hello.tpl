<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title> hello </title>
</head>
<body>
<p>u1</p>
<p>hello： {{.u1.Name}}</p>
<p>年龄： {{.u1.Age}}</p>
<p>性别： {{.u1.Gender}}</p>
<p>m1</p>
<p>hello： {{.m1.name}}</p>
<p>年龄： {{.m1.age}}</p>
<p>性别： {{.m1.gender}}</p>
<hr>
{{range $idx,$hobby := .hobby }}
    <p>i{{$idx}} - {{$hobby}}</p>
{{else}}
    没啥爱好
{{end}}

</body>
</html>