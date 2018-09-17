#!/usr/bin python
import sys
reload(sys)
sys.setdefaultencoding("utf8")
import json
import jieba


def get_configs(path):
    with open(path, "r") as file:
        content = "".join(file.readlines())
        file.close()
    data = json.loads(content)
    return data

def handle(path="", classname=""):
    if classname == "":
        return
    materiel = str(path) + "materiel_" + str(classname) + ".txt"
    with open(materiel, "r") as file:
        lines = file.readlines()
        file.close()
    words = set()
    for line in lines:
        line = line.strip("\n").split("\t")[-1]
        tempwords = jieba.cut(line, cut_all=False)
        for tempword in tempwords:
            words.add(tempword)
    source = str(path) + str(classname) + str(".txt")
    with open(source, "a") as file:
        file.writelines([str(word)+"\n" for word in words])
        file.close()


if __name__ == "__main__":
    path = "config.json"
    classnames = get_configs(path)['CLASSES']
    for classname in classnames:
        handle("", classname)
