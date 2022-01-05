#!/usr/bin/env python3
# -*- coding: utf-8 -*-

'a cp module'

__author__ = 'FunnyDay'

import os
import re


# 获取协议文件
def getCPFile():
    cwd = os.getcwd()
    with open(cwd + '\\client\\client.proto', 'rb') as f:
        return f.read().decode()


# 写入协议文件
def writeCPFile(msg):
    cwd = os.getcwd()
    with open(cwd + '\\client\\client.proto', 'a+', encoding='utf-8') as f:
        return f.write(msg)


content = getCPFile()


# 正则匹配获取命令列表
def regMatch(content):
    s = r'((\w+_)+(Req|Resp|Push))\s*=\s\d+;\s*//(\S*)'
    m = re.findall(s, content, re.M | re.S)
    return m


# 正则匹配获取命令列表
def regMatchTest(content):
    s = r'((\w+_)+(Req|Resp|Push))\s*=\s\d+;\s*//(\S*)'
    m = re.findall(s, content, re.M | re.S)
    for item in m:
        print(item)


# 生成一个message
def newMessage(name, comment):
    name = name.replace('_', '')
    moban = '''
//%s
message %s {
}
    '''
    s = moban % (comment, name)
    return s


# 写入proto
def writeToFile(name, comment):
    name = name.replace('_', '')
    if content.find(name) > -1:
        return
    mes = newMessage(name, comment)
    # mes = mes.encode(encoding="utf-8", errors="strict")
    writeCPFile(mes)


# 生成 message
def generate():
    cps = regMatch(content)
    for item in cps:
        writeToFile(item[0], item[3])


if __name__ == '__main__':
    # regMatchTest(content)
    generate()
