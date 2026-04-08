# Возвращает список узлов
```
f() = 0
f(a) = a
f(a,b) = a b +

f(1)
```
### Для такой программы:
- FunDef(f, {}, Num(0))
- FunDef(f, {a}, Id(a))
- FunDef(f, {a,b}, Plus(Id(a), Id(b)))
- FunInv(FunId(f), args(Num(1)))