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
var toprotopath = "../DiceCC/protobuf/"
var tobinderpath = "../DiceCC/assets/scripts/net/pb/"
var allprotocode = "";

var root;

function main() {
    allprotocode = 'syntax = "proto3";\n//package \n'
    root = __dirname.substring(0, __dirname.indexOf("allproto"));
    if (toprotopath.indexOf("..") == 0) {
        toprotopath = path.join(root, toprotopath);
        console.log(toprotopath);
    }
    readDir(root);
    for (var i = 0; i < replacekey.length; i++) {
        allprotocode = allprotocode.replace(eval("/" + replacekey[i] + "/g"), replacevalue[i]);
    }

    fs.writeFile(toprotopath + 'pb.proto', allprotocode, function (err) {
        if (err) throw err;
        else console.log("写入结束" + toprotopath + 'pb.proto');
    });
     writePbBinder();
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

function writePbBinder() {
    let start = allprotocode.indexOf("enum CP {");
    let end = allprotocode.indexOf("}", start);
    let cpstr = allprotocode.substring(start, end);
    const allcmdid = cpstr.split("\n");
    let encodestr = "this.encodePB = {};\n"
    let decodestr = "this.decodePB = {};\n"
    for (let i = 0; i < allcmdid.length; i++) {
        let cmdid = allcmdid[i];
        let endindex = cmdid.indexOf("=");
        if (endindex > -1) {

            cmdid = cmdid.substring(0, endindex).trim();
            let mess = cmdid.replace(eval("/_/g"), "");
            let hascmd = allprotocode.indexOf(mess) != -1;
            if (cmdid.indexOf("Req") != -1) {
                if (hascmd) {
                    encodestr += "        this.encodePB[pb.CP." + cmdid + "] = pb.encode" + mess + ";\n"
                }
            } else if (cmdid.indexOf("Resp") != -1 || cmdid.indexOf("Push") != -1) {
                if (hascmd)
                    decodestr += "        this.decodePB[pb.CP." + cmdid + "] = pb.decode" + mess + ";\n"
            }
        }
    }
    // console.log(encodestr)
    // console.log(decodestr)
    start = allprotocode.indexOf("enum ErrorType {");
    end = allprotocode.indexOf("}", start);
    cpstr = allprotocode.substring(start, end);
    const allerror = cpstr.split("\n");
    let errorstr = "this.errorMess = {};\n"
    for (let i = 0; i < allerror.length; i++) {
        let errid = allerror[i];
        let endindex = errid.indexOf("=");
        let startindex = errid.indexOf("//");
        if (endindex > -1 && startindex > -1) {
            err = errid.substring(0, endindex).trim();
            mess = errid.substring(startindex + 2, errid.length).trim();
            errorstr += "        this.errorMess[pb.ErrorType." + err + "] = \"" + mess + "\";\n"
        }
    }
    // console.log(errorstr);
    if (tobinderpath.indexOf("..") == 0) {
        tobinderpath = path.join(root, tobinderpath);
    }
    let PbBinder = fs.readFileSync(tobinderpath + 'PbBinder.ts', 'utf8');
    start = PbBinder.indexOf("this.encodePB = {};");
    end = PbBinder.indexOf("}", start + 30);
    oldstr = PbBinder.substring(start, end - 1);
    // console.log(oldstr, encodestr)
    PbBinder = PbBinder.replace(oldstr, encodestr);
    start = PbBinder.indexOf("this.decodePB = {};");
    end = PbBinder.indexOf("}", start + 30);
    oldstr = PbBinder.substring(start, end - 1);
    // console.log(oldstr, decodestr)
    PbBinder = PbBinder.replace(oldstr, decodestr);
    start = PbBinder.indexOf("this.errorMess = {};");
    end = PbBinder.indexOf("}", start + 30);
    oldstr = PbBinder.substring(start, end - 1);
    // console.log(oldstr, errorstr)
    PbBinder = PbBinder.replace(oldstr, errorstr);
    fs.writeFile(tobinderpath + 'PbBinder.ts', PbBinder, function (err) {
        if (err) throw err;
        else console.log("写入结束" + tobinderpath + 'PbBinder.ts');
    });
}
main();