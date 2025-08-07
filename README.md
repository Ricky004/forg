# Forga - File Organizer CLI Tool

A powerful Command Line Interface (CLI) tool written in Go to organize, search, and manage files in directories based on file type, date, or custom rules.

## Features

- **File Organization by Type**: Automatically categorize files into folders based on their extensions (Images, Documents, Videos, Music, etc.)
- **Date-based Organization**: Organize files by their modification date into year/month folder structure
- **Duplicate Detection**: Find and handle duplicate files by removing or relocating them
- **Bulk Renaming**: Add prefixes, suffixes, or sequential numbers to multiple files
- **Advanced File Search**: Search and filter files by name, extension, size, and modification date
- **Custom Configuration**: Use YAML configuration files to define custom file categories
- **Concurrent Processing**: Multi-threaded file operations for better performance

## Installation

### With Golang
**1. Prerequisites**: Ensure you have Go 1.22.0 or higher
 
**2. Run the CLI**: Execute the following command in your terminal:

```bash 
go install github.com/Ricky004/forg@v0.1.2
```
Then, run the CLI from anywhere
```bash
forg
```

### Build from Source
```bash
git clone https://github.com/Ricky004/forg.git
cd forg
go build -o forg
```

## Usage

### Basic Commands

#### Organize Files by Type
```bash
# Organize files in current directory using default categories
forg organize --dir /path/to/directory

# Use custom configuration file
forg organize --dir /path/to/directory --config config.yaml
```

#### Organize Files by Date
```bash
# Organize files by modification date (YYYY/MM structure)
forg organize --dir /path/to/directory --date
```

#### Bulk Renaming
```bash
# Add prefix and suffix to all files
forg organize --dir /path/to/directory --prefix "IMG_" --suffix "_backup"

# Add sequential numbers starting from 100
forg organize --dir /path/to/directory --start-number 100
```

#### Handle Duplicates
```bash
# Remove duplicate files
forg organize --dir /path/to/directory --remove

# Move duplicates to another directory
forg organize --dir /path/to/directory --relocate /path/to/duplicates
```

#### Search Files
```bash
# Search by file name
forg search /path/to/directory --name "document"

# Search by extension
forg search /path/to/directory --extension ".pdf"

# Search by size range
forg search /path/to/directory --min-size "1mb" --max-size "100mb"

# Search by date range
forg search /path/to/directory --after "2024-01-01" --before "2024-12-31"

# Combine multiple filters
forg search /path/to/directory --name "report" --extension ".pdf" --min-size "500kb"
```

## Configuration

### Default File Categories

The tool comes with predefined categories:

- **Images**: `.jpg`, `.jpeg`, `.png`, `.gif`
- **Documents**: `.pdf`, `.doc`, `.docx`, `.txt`
- **Videos**: `.mp4`, `.avi`, `.mkv`
- **Music**: `.mp3`, `.wav`, `.aac`
- **HTML**: `.html`, `.htm`

### Custom Configuration

Create a `config.yaml` file to define custom categories:

```yaml
categories:
  Images:
    - .jpg
    - .jpeg
    - .png
    - .gif
    - .bmp
    - .tiff
  Documents:
    - .pdf
    - .doc
    - .docx
    - .txt
    - .rtf
    - .odt
  Videos:
    - .mp4
    - .avi
    - .mkv
    - .mov
    - .wmv
  Music:
    - .mp3
    - .wav
    - .aac
    - .flac
    - .ogg
  Archives:
    - .zip
    - .rar
    - .7z
    - .tar
    - .gz
```

## Command Reference

### Global Flags
- `--help`: Show help information

### Organize Command
```bash
forg organize [flags]
```

**Flags:**
- `--dir`: Directory to organize (required)
- `--config, -c`: Path to configuration file (optional)
- `--date, -d`: Organize by date
- `--prefix, -p`: Prefix to add to file names
- `--suffix, -s`: Suffix to add to file names
- `--start-number, -n`: Starting number for sequential renaming
- `--remove`: Remove duplicate files
- `--relocate`: Directory to relocate duplicate files

### Search Command
```bash
forg search [directories...] [flags]
```

**Flags:**
- `--name, -n`: Filter by file name
- `--extension, -e`: Filter by file extension
- `--min-size`: Minimum file size (e.g., 3b, 50kb, 100mb)
- `--max-size`: Maximum file size (e.g., 3b, 50kb, 100mb)
- `--before`: Filter files modified before date (YYYY-MM-DD)
- `--after`: Filter files modified after date (YYYY-MM-DD)

## Examples

### Organize Downloads Folder
```bash
# Organize downloads by file type
forg organize --dir ~/Downloads

# Organize downloads by date and add prefix to all files
forg organize --dir ~/Downloads --date --prefix "downloaded_"
```

### Clean Up Photo Directory
```bash
# Organize photos and remove duplicates
forg organize --dir ~/Photos --remove

# Organize photos and move duplicates to backup folder
forg organize --dir ~/Photos --relocate ~/Photos/duplicates
```

### Find Large Files
```bash
# Find files larger than 100MB
forg search ~/Documents --min-size "100mb"

# Find old large video files
forg search ~/Videos --extension ".mp4" --before "2023-01-01" --min-size "1gb"
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [YAML v3](https://gopkg.in/yaml.v3) - YAML parsing

## Author

Created by [Tridip Dam](https://github.com/Ricky004)