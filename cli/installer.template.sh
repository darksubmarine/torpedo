#!/bin/bash

echo " "
echo "* Torpedo Tool Installer"
echo "Commit version: AUTO_REPLACE_COMMIT_VERSION"
echo "Build version: AUTO_REPLACE_BUILD_VERSION"
echo " "

if [ "$1" == "--version" ] || [ "$1" == "-v" ]; then
   exit 0
fi

# Create destination folder
DESTINATION="$HOME/.torpedo"
mkdir -p ${DESTINATION}

# Find __ARCHIVE__ maker, read archive content and decompress it
ARCHIVE=$(awk '/^__ARCHIVE__/ {print NR + 1; exit 0; }' "${0}")
tail -n+${ARCHIVE} "${0}" | tar xpJv -C ${DESTINATION}

ln -sf ${DESTINATION}/bin/AUTO_REPLACE_BINARY_FILENAME ${DESTINATION}/bin/torpedo

echo " "
echo "Installation complete at $DESTINATION"
echo " "
echo "Please add the next line to your .bash_profile file"
echo '    export PATH="$HOME/.torpedo/bin:$PATH"'
echo " "
echo "And then execute the command"
echo '    source $HOME/.bash_profile'

# Exit from the script with success (0)
exit 0

__ARCHIVE__
