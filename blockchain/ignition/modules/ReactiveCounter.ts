import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

export default buildModule("ReactiveCounterModule", (m) => {
  const counter = m.contract("ReactiveCounter");

  m.call(counter, "incBy", [5n]);

  return { counter };
});