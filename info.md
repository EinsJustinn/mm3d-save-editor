
# Majora's Mask 3D Offsets

## Player Info

| Offset            | Description             | Notes / Values                                                               |
|-------------------|--------------------------|-------------------------------------------------------------------------------|
| **0x128-0x136**   | Player Name              | String (up to 8 characters. splitted with spaces)                            |
| **0x140**         | Rupees                   | Max: 999                                                                     |
| **0x1230-0x1231** | Bank Rupees              | Max: 65535                                                                   |
| **0x13A-0x13B**   | Heart Containers         | Number of full heart containers                                              |
| **0x13C-0x13D**   | Current Hearts           | Current heart value                                                          |

## Game Mechanics

| Offset            | Description             | Notes / Values                                                               |
|-------------------|--------------------------|-------------------------------------------------------------------------------|
| **0x10-0x13**     | Time Speed               | `0x00`: Normal<br>`0xFF`: Slow                                               |
| **0x26**          | Current Character        | `0x00`: Fierce Deity<br>`0x01`: Goron<br>`0x02`: Zora<br>`0x03`: Deku<br>`0x04`: Link |

## Save Data Integrity

| Offset            | Description             | Notes / Values                                                               |
|-------------------|--------------------------|-------------------------------------------------------------------------------|
| **0x1A88-0x1A89** | Checksum (CRC-16/ARC)   | Replace with `0D 44`, calculate checksum, then store in reverse order        |

## Item Amount

*for all max 99*
| Offset    | Item                |
|-----------|---------------------|
| 0x1B9     | Arrows              |
| 0x1BE     | Bombs               |
| 0x1BF     | Bombchu             |
| 0x1C0     | Deku Stick          |
| 0x1C1     | Deku Nut            |
| 0x1C2     | Magic Beans         |
| 0x1C4     | Powder Keg          |

## Songs

*for all songs*
| Offset          | Description            | Value |
|-----------------|------------------------|--------|
| **0x1F9**       | Songs                  | `0xF7` |
| **0x1FA**       | Songs                  | `0xCF` |
| **0x1FB**       | Songs                  | `0x01` |

## Teleport Statues (0x14C-0x14D)

*Values are binary coded and can be summed:*

| Value | Location             |
|--------|----------------------|
| 1      | Great Bay Coast      |
| 2      | Zora Cape            |
| 4      | Snowhead             |
| 8      | Mountain Village     |
| 16     | Clock Town           |
| 32     | Milk Road            |
| 64     | Woodfall             |
| 128    | Southern Swamp       |
| 256    | Ikana Canyon         |
| 512    | Stone Tower          |

## Items (Slot Values)

| Item                    | Value |
|--------------------------|--------|
| Hero's Bow               | 0x01   |
| Fire Arrow               | 0x02   |
| Ice Arrow                | 0x03   |
| Light Arrow              | 0x04   |
| Bomb                     | 0x06   |
| Bombchu                  | 0x07   |
| Deku Stick               | 0x08   |
| Deku Nut                 | 0x09   |
| Magic Beans              | 0x0A   |
| Powder Keg               | 0x0C   |
| Lens of Truth            | 0x0E   |
| Hookshot                 | 0x0F   |
| Great Fairy's Sword      | 0x10   |
| Green Potion             | 0x30   |
| Red Potion               | 0x31   |
| Fairy                    | 0x32   |
| Blue Potion              | 0x33   |
| Chateau Romani           | 0x34   |
| Milk                     | 0x35   |
| Fish                     | 0x36   |

## Masks (Slot Values)

| Mask                    | Value |
|--------------------------|--------|
| Postman's Hat            | 0x00   |
| All-Night Mask           | 0x01   |
| Blast Mask               | 0x02   |
| Stone Mask               | 0x03   |
| Great Fairy's Mask       | 0x04   |
| Deku Mask                | 0x05   |
| Keaton Mask              | 0x06   |
| Bremen Mask              | 0x07   |
| Bunny Hood               | 0x08   |
| Don Gero's Mask          | 0x09   |
| Mask of Scents           | 0x0A   |
| Goron Mask               | 0x0B   |
| Romani's Mask            | 0x0C   |
| Troupe Leader's Mask     | 0x0D   |
| Kamaro's Mask            | 0x0D   |
| Kafei's Mask             | 0x0E   |
| Couple's Mask            | 0x0F   |
| Mask of Truth            | 0x10   |
| Zora Mask                | 0x11   |
| Gibdo Mask               | 0x13   |
| Garo's Mask              | 0x14   |
| Captain's Hat            | 0x15   |
| Giant's Mask             | 0x16   |
| Fierce Deity's Mask      | 0x17   |

## Boss Masks

| Offset          | Description            | Value |
|-----------------|------------------------|--------|
| **0x1F8**       | Boss Masks (all)      | `0xCF` |

## Slots

| Offsets             | Description     |
|-------------------|------------------|
| 0x1A30–0x1A43     | Item Slots       |
| 0x1A44–0x1A5B     | Mask Slots       |

## Hotbar Assignments

| Slot | Primary Offset | Secondary Offset |
|------|----------------|------------------|
| X    | 0x154          | 0x168            |
| Y    | 0x153          | 0x167            |
| I    | 0x155          | 0x169            |
| II   | 0x157          | 0x16A            |

**Empty Slot:** `0xFF`
