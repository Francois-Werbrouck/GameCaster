## Game

### Needs:

- Allow Users to be either a PLAYER or a GAMEMASTER
- Allow the PLAYER to create and manage its own PERSO which includes a picture and a HP system.
- The GM can create MONSTERS archetype.
- The GM can create NPC
- The GM can initiate a COMBAT which opens a session for the PLAYER to join in with their PERSO
- The GM can decide which monster achetype will be part of the combat
- All the PLAYER needs to input their initiative for a COMBAT. The monster initiative is automatically rolled
- A COMBAT state must be saved in case a session finished before a COMBAT resume
- We might not save each interaction within the app, so for now, Players and GM needs to be able to manage the HP of each of their PERSOs at anytime during the combat session
- The Combat Subscriber will have their sheets updated at the end of each Subscriber turn.
- The GM needs to be able to add monster mid combat and chose wether or not he wants to manually input the new monster initiative


```mermaid
erDiagram
    USER {
        id int
        email string "Unique"
        password hash
        role type "PLAYER | GAMEMASTER" 
    }
    PERSO {
        id int
        name string
        total_hp int
        picture BLOB
    }
    NPC_ARCHETYPE {
        id int
        name string
        default_hp int
        picture BLOB
    }
    COMBAT {
        id int
        boolean is_active
    }
    COMBATANT{
        id int
        is_perso boolean
        initiative boolean
        entity_id int
        current_hp int
        combat_id int
    }
    USER ||--o{ PERSO : "creates/manages"
    COMBATANT }|--|| NPC_ARCHETYPE : "is"
    USER ||--o{ NPC_ARCHETYPE : "creates"
    USER ||--|| COMBAT : "initiates/joins"
    COMBATANT }|--|| PERSO : "is"
    COMBAT |o--o{ COMBATANT : "includes"
```
