//合并所有proto文件给客户端用
//执行方式  node allproto/allproto

var fs = require('fs');
var path = require('path');
//需要合并的proto文件
var allfilenames = ["client.proto", "common.proto", "error.proto"]
//需要替换的文本
var replacekey = ["common.", "errcode.", "\\[\\(validate.rules\\)"]
var replacevalue = ["", "", ";/////[(validate.rules)"]
//导出文件的路径
var topath = "../DiceCC/protobuf/"
var allprotocode = "";

var root;

function main() {
    allprotocode = 'syntax = "proto3";\n//package \n'
    root = __dirname.substring(0, __dirname.indexOf("allproto"));
    if (topath.indexOf("..") == 0) {
        topath = path.join(root, topath);
        console.log(topath);
    }
    readDir(root);
    for (var i = 0; i < replacekey.length; i++) {
        allprotocode = allprotocode.replace(eval("/" + replacekey[i] + "/g"), replacevalue[i]);
    }

    fs.writeFile(topath + 'pb.proto', allprotocode, function (err) {
        if (err) throw err;
    });
    console.log("压缩结束" + topath + 'pb.proto');

}
//文件遍历方法
function readDir(filePath) {
    //根据文件路径读取文件，返回文件列表
    let files = fs.readdirSync(filePath);
    //遍历读取到的文件列表
    files.forEach(function (filename) {
        //获取当前文件的绝对路径
        var filedir = path.join(filePath, filename);
        //根据文件路径获取文件信息，返回一个fs.Stats对象
        readFile(filedir);
    });
}

function readFile(dir) {
    let stats = fs.statSync(dir)
    var isFile = stats.isFile(); //是文件
    var isDir = stats.isDirectory(); //是文件夹
    if (isFile && isproto(dir)) {
        readProto(dir);
    }
    if (isDir) {
        readDir(dir); //递归，如果是文件夹，就继续遍历该文件夹下面的文件
    }
}

function isproto(dir) {
    if (dir.indexOf(".proto" != -1)) {
        for (var i = 0; i < allfilenames.length; i++) {
            if (dir.indexOf(allfilenames[i]) != -1) {
                return true;
            }
        }
    }
    return false;
}

function readProto(fPath) {
    let filename = fPath.replace(root, "");
    console.log(filename);
    codeStr = fs.readFileSync(fPath, 'utf8');
    let starindex = codeStr.indexOf("option go_package");
    starindex = codeStr.indexOf("\n", starindex) + 1;
    allprotocode += "////" + filename + "\n" + codeStr.substring(starindex);
}
main();