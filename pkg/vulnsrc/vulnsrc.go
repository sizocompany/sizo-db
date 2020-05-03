package vulnsrc

import (
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/alpine"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/amazon"
	archlinux "github.com/sizocompany/sizo-db/pkg/vulnsrc/arch-linux"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/bundler"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/cargo"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/composer"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/debian"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/ghsa"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/glad"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/govulndb"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/node"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/nvd"
	oracleoval "github.com/sizocompany/sizo-db/pkg/vulnsrc/oracle-oval"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/photon"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/python"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/redhat"
	redhatoval "github.com/sizocompany/sizo-db/pkg/vulnsrc/redhat-oval"
	susecvrf "github.com/sizocompany/sizo-db/pkg/vulnsrc/suse-cvrf"
	"github.com/sizocompany/sizo-db/pkg/vulnsrc/ubuntu"
)

type VulnSrc interface {
	Name() string
	Update(dir string) (err error)
}

var (
	// All holds all data sources
	All = []VulnSrc{
		// NVD
		nvd.NewVulnSrc(),

		// OS packages
		alpine.NewVulnSrc(),
		archlinux.NewVulnSrc(),
		redhat.NewVulnSrc(),
		redhatoval.NewVulnSrc(),
		debian.NewVulnSrc(),
		ubuntu.NewVulnSrc(),
		amazon.NewVulnSrc(),
		oracleoval.NewVulnSrc(),
		susecvrf.NewVulnSrc(susecvrf.SUSEEnterpriseLinux),
		susecvrf.NewVulnSrc(susecvrf.OpenSUSE),
		photon.NewVulnSrc(),

		// Language-specific packages
		bundler.NewVulnSrc(),
		composer.NewVulnSrc(),
		node.NewVulnSrc(),
		python.NewVulnSrc(),
		cargo.NewVulnSrc(),
		ghsa.NewVulnSrc(ghsa.Composer),
		ghsa.NewVulnSrc(ghsa.Maven),
		ghsa.NewVulnSrc(ghsa.Npm),
		ghsa.NewVulnSrc(ghsa.Nuget),
		ghsa.NewVulnSrc(ghsa.Pip),
		ghsa.NewVulnSrc(ghsa.Rubygems),
		glad.NewVulnSrc(),
		govulndb.NewVulnSrc(),
	}
)
