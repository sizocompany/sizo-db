package vulnsrc

import (
	"github.com/sizocompany/sizo-db/pkg/types"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/alma"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/alpine"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/amazon"
	archlinux "github.com/sizocompany/sizo-db/pkg/vulnsrc/arch-linux"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/bundler"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/composer"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/debian"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/ghsa"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/glad"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/govulndb"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/mariner"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/node"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/nvd"
	oracleoval "github.com/sizocompany/sizo-db/pkg/vulnsrc/oracle-oval"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/osv"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/photon"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/redhat"
	redhatoval "github.com/sizocompany/sizo-db/pkg/vulnsrc/redhat-oval"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/rocky"
	susecvrf "github.com/sizocompany/sizo-db/pkg/vulnsrc/suse-cvrf"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/ubuntu"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/wolfi"
)

type VulnSrc interface {
	Name() types.SourceID
	Update(dir string) (err error)
}

var (
	// All holds all data sources
	All = []VulnSrc{
		// NVD
		nvd.NewVulnSrc(),

		// OS packages
		alma.NewVulnSrc(),
		alpine.NewVulnSrc(),
		archlinux.NewVulnSrc(),
		redhat.NewVulnSrc(),
		redhatoval.NewVulnSrc(),
		debian.NewVulnSrc(),
		ubuntu.NewVulnSrc(),
		amazon.NewVulnSrc(),
		oracleoval.NewVulnSrc(),
		rocky.NewVulnSrc(),
		susecvrf.NewVulnSrc(susecvrf.SUSEEnterpriseLinux),
		susecvrf.NewVulnSrc(susecvrf.OpenSUSE),
		photon.NewVulnSrc(),
		mariner.NewVulnSrc(),
		wolfi.NewVulnSrc(),

		// Language-specific packages
		bundler.NewVulnSrc(),
		composer.NewVulnSrc(),
		node.NewVulnSrc(),
		ghsa.NewVulnSrc(),
		glad.NewVulnSrc(),
		govulndb.NewVulnSrc(),
		osv.NewVulnSrc(),
	}
)
