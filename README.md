**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-04-25 14:05:56 UTC（2026-04-25 22:05:56 UTC+8）

**代理总数：43**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 43 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 43 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 367ms | ✓ 1224ms | ✓ 779ms | ✓ 954ms | ✓ 1767ms | http |
| 20.18.193.135:8561 | ✓ 1172ms | ✓ 1126ms | ✓ 708ms | ✓ 1071ms | ✓ 837ms | http |
| 20.210.39.153:8561 | ✓ 1157ms | ✓ 1546ms | ✓ 658ms | ✓ 984ms | ✓ 780ms | http |
| 20.78.118.91:8561 | ✓ 1165ms | ✓ 1223ms | ✓ 843ms | ✓ 1080ms | 否 | http |
| 20.78.26.206:8561 | ✓ 1174ms | ✓ 1859ms | ✓ 642ms | ✓ 1029ms | ✓ 769ms | http |
| 168.110.52.228:3128 | ✓ 1768ms | 否 | ✓ 933ms | ✓ 1389ms | ✓ 1002ms | http |
| 206.206.126.177:2412 | ✓ 889ms | 否 | ✓ 1440ms | ✓ 1225ms | 否 | http |
| 80.92.204.47:1081 | 否 | ✓ 1992ms | ✓ 800ms | ✓ 1680ms | ✓ 1332ms | http |
| 43.133.90.161:8888 | 否 | 否 | ✓ 1881ms | ✓ 1056ms | ✓ 1960ms | http |
| 20.27.14.220:8561 | ✓ 1228ms | ✓ 1193ms | ✓ 650ms | 否 | 否 | http |
| 20.210.76.104:8561 | ✓ 1233ms | ✓ 1079ms | ✓ 813ms | 否 | 否 | http |
| 113.160.132.26:8080 | ✓ 1971ms | ✓ 1884ms | 否 | ✓ 1764ms | 否 | http |
| 194.31.87.77:3128 | ✓ 1053ms | ✓ 1918ms | ✓ 1724ms | ✓ 1993ms | 否 | http |
| 20.127.128.70:8080 | ✓ 534ms | 否 | ✓ 1932ms | ✓ 1736ms | ✓ 1456ms | http |
| 45.140.147.82:1081 | ✓ 451ms | 否 | ✓ 491ms | ✓ 1417ms | ✓ 1310ms | http |
| 45.140.147.82:1082 | ✓ 466ms | ✓ 1263ms | ✓ 1179ms | 否 | ✓ 999ms | http |
| 20.27.13.35:8561 | ✓ 1532ms | ✓ 1197ms | ✓ 818ms | ✓ 1089ms | ✓ 930ms | http |
| 20.27.11.248:8561 | ✓ 1613ms | ✓ 1205ms | ✓ 818ms | ✓ 1051ms | ✓ 955ms | http |
| 20.27.15.111:8561 | ✓ 1623ms | ✓ 1802ms | ✓ 676ms | ✓ 1039ms | ✓ 796ms | http |
| 130.61.139.145:3128 | ✓ 519ms | 否 | 否 | ✓ 1508ms | ✓ 1449ms | http |
| 35.225.22.61:80 | ✓ 369ms | ✓ 1555ms | ✓ 721ms | 否 | 否 | http |
| 2.27.54.161:1080 | ✓ 967ms | 否 | ✓ 1723ms | 否 | ✓ 1417ms | http |
| 177.93.132.244:3128 | ✓ 770ms | 否 | ✓ 1013ms | 否 | ✓ 1686ms | http |
| 51.159.125.63:8080 | 否 | ✓ 1410ms | 否 | ✓ 1967ms | ✓ 1524ms | http |
| 47.84.76.30:1080 | ✓ 1034ms | 否 | ✓ 893ms | ✓ 1292ms | 否 | http |
| 47.101.159.19:8899 | ✓ 1097ms | ✓ 1295ms | ✓ 1051ms | ✓ 1304ms | ✓ 1053ms | http |
| 121.230.9.209:1080 | ✓ 1865ms | ✓ 1578ms | ✓ 1494ms | ✓ 1793ms | 否 | http |
| 124.243.150.41:3128 | ✓ 1328ms | ✓ 1635ms | ✓ 1144ms | ✓ 1256ms | ✓ 1265ms | http |
| 94.131.106.231:1081 | ✓ 556ms | ✓ 1310ms | ✓ 462ms | ✓ 1777ms | ✓ 1781ms | http |
| 43.157.41.157:3128 | ✓ 862ms | ✓ 1492ms | ✓ 1054ms | ✓ 1527ms | ✓ 1317ms | http |
| 84.47.150.125:1080 | 否 | 否 | ✓ 1736ms | ✓ 1994ms | ✓ 1608ms | http |
| 201.219.22.2:3128 | ✓ 1410ms | ✓ 1858ms | ✓ 1564ms | ✓ 1745ms | ✓ 1502ms | http |
| 1.231.81.166:3128 | ✓ 1505ms | 否 | ✓ 1152ms | ✓ 1236ms | ✓ 1040ms | http |
| 62.113.119.14:8080 | ✓ 1762ms | 否 | ✓ 1025ms | ✓ 1856ms | ✓ 1428ms | http |
| 68.183.199.89:1080 | ✓ 1335ms | 否 | 否 | ✓ 985ms | ✓ 1140ms | http |
| 45.76.207.177:40000 | ✓ 1524ms | 否 | ✓ 1788ms | ✓ 1340ms | ✓ 1069ms | http |
| 152.70.91.193:40000 | ✓ 1531ms | 否 | 否 | ✓ 1964ms | ✓ 1975ms | http |
| 47.84.59.16:1080 | ✓ 1525ms | ✓ 1902ms | ✓ 1085ms | ✓ 1292ms | ✓ 1022ms | http |
| 8.219.97.248:80 | ✓ 933ms | 否 | ✓ 1464ms | 否 | ✓ 1527ms | http |
| 45.153.231.229:8080 | ✓ 813ms | 否 | ✓ 1159ms | 否 | ✓ 1739ms | http |
| 117.236.124.166:3128 | ✓ 1017ms | 否 | ✓ 1107ms | ✓ 1988ms | 否 | http |
| 61.52.131.172:8443 | ✓ 1076ms | ✓ 1370ms | ✓ 1271ms | ✓ 1493ms | ✓ 1131ms | http |
| 183.232.248.73:7890 | ✓ 1107ms | ✓ 1356ms | ✓ 1102ms | ✓ 1310ms | ✓ 1042ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
