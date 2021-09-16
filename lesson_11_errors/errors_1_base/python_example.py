"""
Если у нас происходит что-то не предусмотренное и программа ломается,
интерпретатор Python останавливает обработку кода, а так же создает специальный объект
Exception, содержащий информацию об ошибке.

Делается это с помощью конструкции try-except.
"""
try:
    print(undefiened_value)
except Exception:
    # Здесь описано, что делать после появления объекта Exception.
    print("Exception catched")


try:
    print(undefiened_value)
except Exception as e:
    # Исключение можно словить, и, как объект, передать дальше и делать с ним все, что нужно.
    print(f"Exception catched: {e}")


try:
    print(undefiened_value)
except Exception as e:
    # Можем увидеть, что исключение - объект определенного класса.
    print(type(e))

"""
Исключение - это такой же класс, как и все другие, и из них,
точно так же, создаются объекты.
Исключений много, под каждый тип ошибок, и они сконструированы с помощью наследования,
от базового класса Exception.

Посмотреть их можно тут:
https://pythonworld.ru/tipy-dannyx-v-python/isklyucheniya-v-python-konstrukciya-try-except-dlya-obrabotki-isklyuchenij.html
"""

try:
    print(undefiened_value)
except NameError as name_error:
    print(f"Wrong value name catched - {name_error}")


try:
    print(undefiened_value)
    245/0
except NameError as name_error:
    print(f"Wrong value name catched - {name_error}")


try:
    245/0
except ZeroDivisionError as zero_division_error:
    print(f"Exception catched - {zero_division_error}")

"""
Исключения можно вызывать самостоятельно, с помощью ключевого слова raise. 
Можно их кастомизировать и создавать самому, с помощью наследования.
"""

raise ZeroDivisionError("ZeroDivisionException catched")
raise Exception("ZeroDivisionException catched")
raise Exception("Exception catched")