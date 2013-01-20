def count_letters(s):
    return len(''.join(s.split()))

letters1 = count_letters("one two three four five six seven eight nine")
letters10 = count_letters("twenty thirty forty fifty sixty seventy eighty ninety")

total_letters = 0

# 10-19
total_letters += count_letters("ten eleven twelve thirteen fourteen fifteen sixteen seventeen eighteen nineteen")

# 1-9, 20-99
total_letters += letters10 * 10
total_letters += letters1 * 9

letters_1_99 = total_letters

# 100-999
total_letters += letters1 * 100
total_letters += len("hundred") * 900
total_letters += len("and") * 99 * 9
total_letters += letters_1_99 * 9

# 1000
total_letters += count_letters("one thousand")

print(total_letters)
