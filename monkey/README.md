Go言語によるインタプリタの実装  
「Go言語でつくるインタプリタ」  

reference : https://interpreterbook.com/  
sorcecode : https://interpreterbook.com/waiig_code_1.4  

---
### 抽象構文木
- 中置識別子の構文解析  
1, その式に必要な要素を構造体に切り出す  
1 + 2 であれば、 expression operator expression が必要  