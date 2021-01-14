export type IconsId =
  | "cargo-ship-50"
  | "cpu"
  | "disk"
  | "kubernetes"
  | "linux"
  | "memory"
  | "triangle-marker"
  | "ubuntu";

export type IconsKey =
  | "CargoShip_50"
  | "Cpu"
  | "Disk"
  | "Kubernetes"
  | "Linux"
  | "Memory"
  | "TriangleMarker"
  | "Ubuntu";

export enum Icons {
  CargoShip_50 = "cargo-ship-50",
  Cpu = "cpu",
  Disk = "disk",
  Kubernetes = "kubernetes",
  Linux = "linux",
  Memory = "memory",
  TriangleMarker = "triangle-marker",
  Ubuntu = "ubuntu",
}

export const ICONS_CODEPOINTS: { [key in Icons]: string } = {
  [Icons.CargoShip_50]: "61697",
  [Icons.Cpu]: "61698",
  [Icons.Disk]: "61699",
  [Icons.Kubernetes]: "61700",
  [Icons.Linux]: "61701",
  [Icons.Memory]: "61702",
  [Icons.TriangleMarker]: "61703",
  [Icons.Ubuntu]: "61704",
};
