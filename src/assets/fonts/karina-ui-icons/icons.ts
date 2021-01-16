export type IconsId =
  | "canary-checker"
  | "cargo-ship-50"
  | "cpu"
  | "disk"
  | "grafana"
  | "karma"
  | "kibana"
  | "kubernetes"
  | "linux"
  | "memory"
  | "prometheus"
  | "triangle-marker"
  | "ubuntu";

export type IconsKey =
  | "CanaryChecker"
  | "CargoShip_50"
  | "Cpu"
  | "Disk"
  | "Grafana"
  | "Karma"
  | "Kibana"
  | "Kubernetes"
  | "Linux"
  | "Memory"
  | "Prometheus"
  | "TriangleMarker"
  | "Ubuntu";

export enum Icons {
  CanaryChecker = "canary-checker",
  CargoShip_50 = "cargo-ship-50",
  Cpu = "cpu",
  Disk = "disk",
  Grafana = "grafana",
  Karma = "karma",
  Kibana = "kibana",
  Kubernetes = "kubernetes",
  Linux = "linux",
  Memory = "memory",
  Prometheus = "prometheus",
  TriangleMarker = "triangle-marker",
  Ubuntu = "ubuntu",
}

export const ICONS_CODEPOINTS: { [key in Icons]: string } = {
  [Icons.CanaryChecker]: "61697",
  [Icons.CargoShip_50]: "61698",
  [Icons.Cpu]: "61699",
  [Icons.Disk]: "61700",
  [Icons.Grafana]: "61701",
  [Icons.Karma]: "61702",
  [Icons.Kibana]: "61703",
  [Icons.Kubernetes]: "61704",
  [Icons.Linux]: "61705",
  [Icons.Memory]: "61706",
  [Icons.Prometheus]: "61707",
  [Icons.TriangleMarker]: "61708",
  [Icons.Ubuntu]: "61709",
};
