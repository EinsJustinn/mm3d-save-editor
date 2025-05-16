# MM3D Save Editor

> [!WARNING]
> Work in progress

This tool allows you to edit save files from **Majora's Mask 3D**, including changes to the player's name, money amount, and bank balance.

## Parameters

| Parameter | Description                                        | Example               |
|-----------|----------------------------------------------------|-----------------------|
| `--file`  | Path to the save file to be edited (required).     | `--file save0.bin`    |
| `--name`  | New player name (maximum 16 characters).           | `--name Link`         |
| `--money` | Amount of available money (between 0 and 999).     | `--money 500`         |
| `--bank`  | Balance of the bank account (between 0 and 65535). | `--bank 10000`        |


## Example

```bash
./mm3d-save-editor --file save0.bin --name Link --money 100 --bank 500
```
