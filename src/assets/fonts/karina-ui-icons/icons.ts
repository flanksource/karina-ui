export type IconsId =
  | "about-50"
  | "approval-50"
  | "birdie"
  | "bitbucket-50-2"
  | "bitbucket-50"
  | "calendar-50"
  | "canary-checker"
  | "cancel-50"
  | "cargo-ship-50"
  | "cluster"
  | "commit"
  | "cpu"
  | "disk"
  | "git"
  | "grafana"
  | "heartline"
  | "idea"
  | "karma"
  | "kibana"
  | "kubernetes"
  | "linux"
  | "memory"
  | "node"
  | "prometheus"
  | "ram"
  | "server"
  | "smile"
  | "storage"
  | "triangle-marker"
  | "ubuntu"
  | "vmware-50"
  | "webhook-50";

export type IconsKey =
  | "About_50"
  | "Approval_50"
  | "Birdie"
  | "Bitbucket_50_2"
  | "Bitbucket_50"
  | "Calendar_50"
  | "CanaryChecker"
  | "Cancel_50"
  | "CargoShip_50"
  | "Cluster"
  | "Commit"
  | "Cpu"
  | "Disk"
  | "Git"
  | "Grafana"
  | "Heartline"
  | "Idea"
  | "Karma"
  | "Kibana"
  | "Kubernetes"
  | "Linux"
  | "Memory"
  | "Node"
  | "Prometheus"
  | "Ram"
  | "Server"
  | "Smile"
  | "Storage"
  | "TriangleMarker"
  | "Ubuntu"
  | "Vmware_50"
  | "Webhook_50";

export enum Icons {
  About_50 = "about-50",
  Approval_50 = "approval-50",
  Birdie = "birdie",
  Bitbucket_50_2 = "bitbucket-50-2",
  Bitbucket_50 = "bitbucket-50",
  Calendar_50 = "calendar-50",
  CanaryChecker = "canary-checker",
  Cancel_50 = "cancel-50",
  CargoShip_50 = "cargo-ship-50",
  Cluster = "cluster",
  Commit = "commit",
  Cpu = "cpu",
  Disk = "disk",
  Git = "git",
  Grafana = "grafana",
  Heartline = "heartline",
  Idea = "idea",
  Karma = "karma",
  Kibana = "kibana",
  Kubernetes = "kubernetes",
  Linux = "linux",
  Memory = "memory",
  Node = "node",
  Prometheus = "prometheus",
  Ram = "ram",
  Server = "server",
  Smile = "smile",
  Storage = "storage",
  TriangleMarker = "triangle-marker",
  Ubuntu = "ubuntu",
  Vmware_50 = "vmware-50",
  Webhook_50 = "webhook-50",
}

export const ICONS_CODEPOINTS: { [key in Icons]: string } = {
  [Icons.About_50]: "61697",
  [Icons.Approval_50]: "61698",
  [Icons.Birdie]: "61699",
  [Icons.Bitbucket_50_2]: "61700",
  [Icons.Bitbucket_50]: "61701",
  [Icons.Calendar_50]: "61702",
  [Icons.CanaryChecker]: "61703",
  [Icons.Cancel_50]: "61704",
  [Icons.CargoShip_50]: "61705",
  [Icons.Cluster]: "61706",
  [Icons.Commit]: "61707",
  [Icons.Cpu]: "61708",
  [Icons.Disk]: "61709",
  [Icons.Git]: "61710",
  [Icons.Grafana]: "61711",
  [Icons.Heartline]: "61712",
  [Icons.Idea]: "61713",
  [Icons.Karma]: "61714",
  [Icons.Kibana]: "61715",
  [Icons.Kubernetes]: "61716",
  [Icons.Linux]: "61717",
  [Icons.Memory]: "61718",
  [Icons.Node]: "61719",
  [Icons.Prometheus]: "61720",
  [Icons.Ram]: "61721",
  [Icons.Server]: "61722",
  [Icons.Smile]: "61723",
  [Icons.Storage]: "61724",
  [Icons.TriangleMarker]: "61725",
  [Icons.Ubuntu]: "61726",
  [Icons.Vmware_50]: "61727",
  [Icons.Webhook_50]: "61728",
};
