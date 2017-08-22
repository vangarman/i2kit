package aws

const elbTemplate = `
  "MyLoadBalancer" : {
    "Type" : "AWS::ElasticLoadBalancing::LoadBalancer",
    "Properties" : {
      "Instances" : [
      ],
      "Listeners" : [
      ],
    }
  }
`

const elbInstanceTemplate = `
  { "Ref" : "logical name of AWS::EC2::Instance resource 1" }
`

const elbListenerTemplate = `
  {
    "LoadBalancerPort" : "80",
    "InstancePort" : "80",
    "Protocol" : "HTTP"
  }
`
