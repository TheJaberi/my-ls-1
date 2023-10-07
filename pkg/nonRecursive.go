package Myls

import "fmt"
func NonRecursive(dir string, flags Flags) {
	// Read the directory
	filesAndFolders, err := Read(dir, flags)
	if err != nil {
		fmt.Println(err)
		Fail = append(Fail, dir)
		return
	}
	// Sort the files and folders
	sortFilesAndFolders(filesAndFolders, flags)
	// Sort by modification time if -t flag is set
	if flags.T {
		sortByModification(filesAndFolders, flags)
	}
	// Print the files and folders
	for _, file := range filesAndFolders {
		printFileOrDir(file, file.Info.IsDir(), flags)
	}
	// fmt.Println()
	for _, file := range filesAndFolders {
		TotalBlocks += file.blockSize
	}
}