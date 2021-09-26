for i in [
    "Scripts/basegenerator.py",
    "Scripts/colourgenerator.py",
    "Scripts/glowgenerator.py",
    "Scripts/grad1generator.py",
    "Scripts/grad2generator.py",
    "Scripts/positiongenerator.py"
]:
    with open(i) as file:
        exec(file.read())