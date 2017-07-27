/*
 * This file is part of arduino-cli.
 *
 * arduino-cli is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 *
 * Copyright 2017 BCMI LABS SA (http://www.arduino.cc/)
 */

package cores

import (
	"github.com/bcmi-labs/arduino-cli/common"
)

const (
	//PackageIndexURL contains the index URL for core packages.
	PackageIndexURL = "http://downloads.arduino.cc/packages/package_index.json"
)

// getDownloadCacheFolder gets the folder where temp installs are stored until installation complete (libraries).
func getDownloadCacheFolder() (string, error) {
	return common.GetDownloadCacheFolder("packages")
}

// DownloadPackagesFile downloads the core packages index file from arduino repository.
func DownloadPackagesFile() error {
	return common.DownloadPackageIndex(IndexPath, PackageIndexURL)
}

// download core
// check core deps, needs statuscontext
// download tools
// download tool -> download Release -> downloadPackage
