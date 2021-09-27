# ! THIS IS DEPRECATED. EDIT BASE.CSS IN THE FILE

# I know it is more efficient putting it all on one line but it is easier to edit this way.
output =  "*{font-family:'Segoe UI',Tahoma,Geneva,Verdana,sans-serif;}"
output += ".center-text{text-align:center;}"
output += ".flex{display:flex;}"
output += ".sticky{position:-webkit-sticky;position:sticky;top:0;}"

output += ".nav{display:flex;flex-wrap:wrap;align-items:center;justify-content:center;}"
output += ".nav-link * {text-decoration:none;}.nav-link {text-decoration:none}"
output += ".nav-title{order:-1;}.nav-title, nav-title * {font-family:'Century Gothic';font-weight:bold;}"

print("Writing to file...")
import os
os.path.join( os.getcwd(), '..', 'CSS/base.css' )
with open(os.path.join( os.getcwd(), "CSS", 'base.css'), "w") as file:
    file.write(output)
    
print("""
   ===========
 ===         ===
===   DONE!   ===
 ===         ===
   ===========
""")