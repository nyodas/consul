class Consul < FPM::Cookery::Recipe
  name 'consul'

  version ENV['PACKAGE_VERSION']? ENV['PACKAGE_VERSION'] : ENV['CI_COMMIT_TAG'][/[0-9\-\._].*/]
  revision '1'
  description 'Consul Service Discovery Platform'

  homepage 'https://www.consul.io'
  # source "https://releases.hashicorp.com/consul/#{version}/consul_#{version}_linux_amd64.zip"
  # sha256 '87c9bce5128654e8f785a324ad88912101acda4c7915c73e74fdbd875d8e7806'
  source "file:///go/src/github.com/hashicorp/consul/pkg/bin/linux_amd64/consul"

  maintainer 'Datadog <package@datadoghq.com>'
  vendor 'datadog'

  license 'Mozilla Public License, version 2.0'

  conflicts 'consul'
  replaces 'consul'

  #post_install 'postinst'

  def build
  end

  def install
    prefix('local/bin').mkdir
    prefix('local/bin').install 'consul'
  end
end
