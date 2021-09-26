for i in [
    "Scripts/basegenerator.py",
    "Scripts/colourgenerator.py",
    "Scripts/glowgenerator.py",
    "Scripts/gradgenerator.py",
    "Scripts/positiongenerator.py"
]:
    with open(i) as file:
        exec(file.read())