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

最后更新：2026-03-08 07:46:02 UTC（2026-03-08 15:46:02 UTC+8）

**代理总数：73**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1324ms | ✓ 1471ms | ✓ 912ms | ✓ 856ms | ✓ 778ms | http |
| 205.209.118.30:3138 | ✓ 953ms | 否 | ✓ 868ms | ✓ 1337ms | ✓ 1222ms | http |
| 121.128.121.54:3128 | ✓ 1665ms | ✓ 1187ms | 否 | 否 | ✓ 1869ms | http |
| 103.84.95.54:7890 | ✓ 736ms | 否 | ✓ 1107ms | ✓ 808ms | ✓ 624ms | http |
| 202.155.12.161:443 | ✓ 1825ms | 否 | 否 | ✓ 1850ms | ✓ 1291ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1092ms | ✓ 1897ms | ✓ 1180ms | 否 | http |
| 112.78.187.186:8080 | ✓ 1759ms | 否 | ✓ 1205ms | ✓ 1329ms | ✓ 1296ms | http |
| 114.4.251.26:8080 | ✓ 1759ms | 否 | ✓ 1171ms | ✓ 1341ms | ✓ 1325ms | http |
| 150.107.140.238:3128 | ✓ 1987ms | 否 | ✓ 1807ms | ✓ 1778ms | ✓ 1071ms | http |
| 120.92.212.16:8890 | ✓ 1905ms | ✓ 1976ms | 否 | 否 | ✓ 1129ms | http |
| 172.212.68.37:3128 | ✓ 1469ms | 否 | ✓ 695ms | ✓ 1540ms | ✓ 1309ms | http |
| 193.168.173.136:443 | ✓ 765ms | ✓ 1996ms | ✓ 977ms | 否 | 否 | http |
| 81.70.169.194:80 | ✓ 1246ms | ✓ 1877ms | ✓ 931ms | 否 | ✓ 1566ms | http |
| 46.183.25.8:443 | ✓ 960ms | 否 | ✓ 1107ms | ✓ 1879ms | 否 | http |
| 152.42.213.210:80 | ✓ 1491ms | 否 | ✓ 783ms | ✓ 1007ms | ✓ 1113ms | http |
| 152.42.213.210:8080 | ✓ 1610ms | 否 | ✓ 709ms | ✓ 1278ms | ✓ 1076ms | http |
| 163.44.126.97:3128 | ✓ 1423ms | 否 | 否 | ✓ 881ms | ✓ 984ms | http |
| 14.56.107.244:3128 | ✓ 1427ms | 否 | 否 | ✓ 1987ms | ✓ 736ms | http |
| 120.92.212.16:7890 | ✓ 1304ms | 否 | ✓ 1562ms | ✓ 1167ms | ✓ 1416ms | http |
| 117.159.239.49:22222 | ✓ 955ms | ✓ 1036ms | 否 | 否 | ✓ 821ms | http |
| 183.249.5.111:22222 | ✓ 706ms | ✓ 1144ms | ✓ 706ms | ✓ 921ms | ✓ 686ms | http |
| 120.240.35.173:22222 | ✓ 875ms | ✓ 1200ms | 否 | ✓ 1129ms | ✓ 1913ms | http |
| 120.240.35.178:22222 | ✓ 807ms | 否 | ✓ 965ms | 否 | ✓ 853ms | http |
| 120.198.141.79:22222 | ✓ 834ms | ✓ 1091ms | ✓ 995ms | ✓ 1111ms | ✓ 862ms | http |
| 120.240.35.177:22222 | ✓ 911ms | ✓ 1220ms | 否 | ✓ 1360ms | 否 | http |
| 120.240.35.160:22222 | ✓ 836ms | ✓ 1160ms | ✓ 1971ms | ✓ 1082ms | ✓ 858ms | http |
| 120.240.35.176:22222 | ✓ 856ms | ✓ 1145ms | ✓ 1596ms | ✓ 1091ms | ✓ 1205ms | http |
| 14.225.211.139:7890 | 否 | 否 | ✓ 859ms | ✓ 1874ms | ✓ 1165ms | http |
| 120.240.35.161:22222 | ✓ 853ms | ✓ 1266ms | ✓ 835ms | ✓ 1034ms | ✓ 864ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1497ms | ✓ 1391ms | ✓ 1898ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1393ms | ✓ 1169ms | ✓ 1307ms | 否 | http |
| 190.9.109.196:999 | 否 | ✓ 1297ms | ✓ 1298ms | ✓ 1345ms | ✓ 1512ms | http |
| 190.9.109.199:999 | 否 | ✓ 1428ms | ✓ 1355ms | ✓ 1428ms | ✓ 1221ms | http |
| 222.184.48.241:22222 | ✓ 1802ms | 否 | 否 | ✓ 1477ms | ✓ 1162ms | http |
| 61.72.110.54:3128 | ✓ 1635ms | 否 | ✓ 1509ms | 否 | ✓ 1880ms | http |
| 39.104.201.40:7890 | ✓ 908ms | ✓ 1955ms | ✓ 1855ms | ✓ 1168ms | ✓ 899ms | http |
| 138.124.53.25:7443 | ✓ 1138ms | 否 | ✓ 1745ms | ✓ 1802ms | ✓ 1426ms | http |
| 193.228.139.78:8888 | ✓ 1141ms | ✓ 1829ms | 否 | 否 | ✓ 1544ms | http |
| 45.140.147.155:1081 | ✓ 1142ms | ✓ 1956ms | 否 | ✓ 1864ms | ✓ 1797ms | http |
| 47.77.193.180:1080 | 否 | 否 | ✓ 433ms | ✓ 711ms | ✓ 537ms | http |
| 38.180.2.107:3128 | ✓ 1027ms | 否 | ✓ 1789ms | 否 | ✓ 1831ms | http |
| 178.236.245.17:3128 | ✓ 1286ms | 否 | ✓ 1862ms | 否 | ✓ 1832ms | http |
| 117.159.239.51:22222 | ✓ 987ms | 否 | 否 | ✓ 1079ms | ✓ 835ms | http |
| 165.227.5.10:8888 | ✓ 1311ms | 否 | ✓ 1051ms | ✓ 1783ms | ✓ 498ms | http |
| 103.139.138.194:3128 | ✓ 1880ms | 否 | ✓ 1213ms | ✓ 1445ms | ✓ 1127ms | http |
| 14.225.222.185:7890 | ✓ 1327ms | 否 | ✓ 1580ms | 否 | ✓ 815ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1104ms | ✓ 1124ms | ✓ 915ms | http |
| 46.249.103.192:443 | ✓ 794ms | 否 | ✓ 1232ms | ✓ 1969ms | 否 | http |
| 45.140.147.82:1081 | ✓ 658ms | ✓ 1616ms | ✓ 1014ms | ✓ 1756ms | ✓ 1271ms | http |
| 61.72.221.194:3128 | ✓ 853ms | ✓ 1088ms | ✓ 822ms | ✓ 1453ms | ✓ 724ms | http |
| 42.115.72.27:2049 | ✓ 1371ms | 否 | ✓ 1439ms | ✓ 1645ms | 否 | http |
| 120.240.29.168:22222 | ✓ 873ms | 否 | ✓ 1214ms | 否 | ✓ 883ms | http |
| 101.230.73.57:29999 | ✓ 774ms | 否 | ✓ 807ms | 否 | ✓ 788ms | http |
| 222.184.48.252:22222 | ✓ 946ms | ✓ 1107ms | 否 | ✓ 1530ms | 否 | http |
| 121.230.8.153:1080 | ✓ 1909ms | ✓ 1331ms | ✓ 1981ms | 否 | ✓ 1006ms | http |
| 45.136.198.40:3128 | ✓ 1362ms | ✓ 1678ms | 否 | ✓ 1739ms | ✓ 1331ms | http |
| 45.129.141.143:3128 | ✓ 990ms | 否 | ✓ 1746ms | 否 | ✓ 1809ms | http |
| 24.199.124.151:3128 | ✓ 221ms | ✓ 875ms | ✓ 908ms | ✓ 649ms | ✓ 713ms | http |
| 212.175.29.184:8080 | ✓ 889ms | ✓ 1835ms | ✓ 1991ms | 否 | 否 | http |
| 101.43.255.96:80 | ✓ 1037ms | ✓ 1188ms | ✓ 1830ms | 否 | 否 | http |
| 47.101.149.27:9010 | ✓ 1236ms | ✓ 1762ms | 否 | 否 | ✓ 1667ms | http |
| 168.235.110.63:3128 | 否 | ✓ 1975ms | ✓ 1224ms | ✓ 1854ms | ✓ 957ms | http |
| 211.217.231.234:8080 | 否 | 否 | ✓ 820ms | ✓ 917ms | ✓ 722ms | http |
| 113.59.32.161:22222 | 否 | 否 | ✓ 1501ms | ✓ 1332ms | ✓ 1033ms | http |
| 62.113.119.14:8080 | ✓ 770ms | 否 | ✓ 1185ms | ✓ 1646ms | ✓ 1232ms | http |
| 106.14.203.63:3333 | 否 | 否 | ✓ 921ms | ✓ 1047ms | ✓ 1893ms | http |
| 85.209.0.118:3128 | ✓ 821ms | 否 | ✓ 1650ms | 否 | ✓ 1915ms | http |
| 103.215.36.88:19115 | ✓ 1155ms | ✓ 1633ms | 否 | 否 | ✓ 1346ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1877ms | 否 | ✓ 1475ms | ✓ 941ms | http |
| 103.39.51.190:8080 | ✓ 1758ms | 否 | 否 | ✓ 1283ms | ✓ 1408ms | http |
| 103.82.23.118:5249 | 否 | 否 | ✓ 1721ms | ✓ 1734ms | ✓ 1546ms | http |
| 103.183.10.203:3125 | 否 | 否 | ✓ 1851ms | ✓ 1868ms | ✓ 1500ms | http |
| 42.115.72.27:2039 | ✓ 1458ms | 否 | 否 | ✓ 1626ms | ✓ 1455ms | http |

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
