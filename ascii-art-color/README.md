START

Read os.Args

IF number of args invalid
    print error
    STOP

Extract color flag
Split flag by "="
Get actual color

IF color not in map
    print error
    STOP

IF args length == 3
    text = input text
    print colored(text)
    STOP

IF args length == 4
    substring = args[2]
    text = args[3]

    result = empty string
    i = 0

    WHILE i < length(text)
        IF substring starts at i
            append colored(substring) to result
            i = i + length(substring)
        ELSE
            append text[i] to result
            i = i + 1
        ENDIF
    ENDWHILE

    print result

END