import os
from openpyxl import load_workbook

class CardSwitcher(object):
    def __init__(self, file_name, sheet_name):
        self.__file_name = file_name
        self.__sheet_name = sheet_name

    def load_excel(self):
        wb = load_workbook(self.__file_name)
        self.sheet = wb[self.__sheet_name]

    def switch():
        return

    def print_info(self):
        print("info : %s" % self.sheet)


if __name__ == '__main__':
    cwd = os.getcwd()
    cs = CardSwitcher(cwd+'\\tool\\king.xlsx', '通用卡牌')
    cs.load_excel()
    cs.print_info()
