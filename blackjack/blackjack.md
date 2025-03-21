https://bicyclecards.com/how-to-play/blackjack

Goal: Have a higher hand value than the dealer and not go beyond 21.
There is only one round in the game of black jack.

Actors
- Dealer
- Players

PlayerActions
- Stand
- Bet
- Hit
- Bust
- NotifyBlackJack
- SplittingPairs

DealerActions

- Deal
- BetDistribution
- Shuffle
- PickCard


Game
has a CardDeck
has a Dealer
has many Players

Game/Dealer is the CPU

Dealing Phase
- Starts off with first round where 2 cards are dealt to all players and the dealer. All players get cards face up, dealer gets it face down.

Betting Phase
- Players place bet

Naturals Evaluation Phase
- If the dealers card is a natural, collect all bets from user who do not have a natural.
- If any player has an natural, its a tie break with the dealer, the player gets his bet.
- The dealer can open the face down card only when the first card is an ace or a ten.

Play
- Players chances go from left to right.
- Players can choose to Stand, Hit, Bust, Double Down, Split Up or Surrender

Giving Away
- If Dealer wins they get all the bet amount.
- If a player wins
Player Action
- Stand: skip the round, not drawing anything
- Hit: when trying to go close to the 21 mark
- Bust: when player goes above 21. 
- Double Down: Increase the Bet by 100%
- Split Up: If the first cards are equal. Play with two hands, by individually drawing and placing bets.
- Surrender: Forfetit get half the bet amount

Dealer cannot double, split or surrender


CardDeck
    - cards list

Card
    State
    - map[string][string]

Balance
    State
    - amount double

Player
    - has cards
    - has balance

    behaviours
    - Actions (Stand/Hit/Bust/Double Down/Split Up/Surrender) 

Dealer/CPU
    - has players
    - has cards
        
    State
    - betAmount
    
    Behaviours
    - DealCards
    - GatherBets
    - EvaluateNormals
    - Play
    - DeclareWinner
    - DistriubteAmounts

Game
    - has a Dealer
    - has many Players
    - has a GameConfig

    Behaviour
    - start, stop, stats, initiateNewRound

GameConfig
    - has many CardDecks
    - has FeatureToggles
    

