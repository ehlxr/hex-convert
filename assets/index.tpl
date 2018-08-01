<h1>转换十进制</h1>
<form action="d" method="post">
    <input type="text" name="scale" onkeyup="this.value=this.value.replace(/\D/g,'')" placeholder="进制数"/><br>
    <input type="text" name="data" placeholder="数值"/>
    <input type="submit" value="转换"/>
</form>

<h1>转换二进制</h1>
<form action="b" method="post">
    <input type="text" name="scale" onkeyup="this.value=this.value.replace(/\D/g,'')" placeholder="进制数"/><br>
    <input type="text" name="data" placeholder="数值"/>
    <input type="submit" value="转换" />
</form>