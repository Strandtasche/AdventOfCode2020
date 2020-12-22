def recursive(deck1, deck2):  #returns tuple of bool, list[int]
    seenCombs = []
    counter = 0
    local_game = next_id()
    # print("new recursion!")
    while deck1 and deck2:  #both are not empty
        if (deck1, deck2) in seenCombs:
            # print("found endless loop!")
            return True, deck1
        else:
            seenCombs.append((deck1.copy(), deck2.copy()))
        counter += 1
        card1 = deck1.pop(0)
        card2 = deck2.pop(0)
        # print(f"round {counter} in game {local_game}")
        if len(deck1) >= card1 and len(deck2) >= card2:
            win1, _ = recursive(deck1[:card1].copy(), deck2[:card2].copy())
            if win1:
                deck1.append(card1)
                deck1.append(card2)
            else:
                deck2.append(card2)
                deck2.append(card1)
        else:
            if card1 > card2:
                deck1.append(card1)
                deck1.append(card2)
            elif card1 < card2:
                deck2.append(card2)
                deck2.append(card1)
            else:
                raise ValueError("broken decks")
    if not deck1:
        # print(f"winnter of game {local_game} is 2")
        return False, deck2
    elif not deck2:
        # print(f"winnter of game {local_game} is 1")
        return True, deck1

def next_id():
    global iid
    res = iid
    iid += 1
    return res


with open('./data/input22.txt') as f:
    decks = [l.rstrip('\n') for l in f.read().split('\n\n')]

deck_player1 = [int(i) for i in decks[0].split('\n')[1:]]
deck_player2 = [int(i) for i in decks[1].split('\n')[1:]]

iid = 1

win1, deck = recursive(deck_player1, deck_player2)

print(f"win1: {win1}")
print(deck)

a = sum([i * j for i,j in zip(range(1, len(deck)+1), deck[::-1])])
# b = sum([i * j for i,j in zip(range(1, len(deck_player2)+1), deck_player2[::-1])])
print(f'p1: {a}')