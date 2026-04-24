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

最后更新：2026-04-24 19:50:03 UTC（2026-04-25 03:50:03 UTC+8）

**代理总数：55**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 55 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 55 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 789ms | 否 | ✓ 1168ms | ✓ 1275ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1641ms | ✓ 976ms | ✓ 964ms | ✓ 927ms | ✓ 733ms | http |
| 113.160.132.26:8080 | ✓ 1414ms | ✓ 1390ms | 否 | 否 | ✓ 996ms | http |
| 45.76.207.177:40000 | ✓ 650ms | 否 | ✓ 1134ms | ✓ 1512ms | ✓ 1040ms | http |
| 62.113.119.14:8080 | ✓ 1277ms | 否 | ✓ 1145ms | ✓ 1702ms | ✓ 1498ms | http |
| 91.233.223.147:3128 | ✓ 1204ms | 否 | ✓ 1213ms | 否 | ✓ 1647ms | http |
| 120.92.108.86:7890 | ✓ 1167ms | 否 | ✓ 1383ms | 否 | ✓ 1793ms | http |
| 47.85.51.197:1080 | ✓ 281ms | 否 | ✓ 816ms | ✓ 1797ms | 否 | http |
| 154.53.61.174:3128 | ✓ 938ms | ✓ 1144ms | ✓ 991ms | ✓ 1530ms | ✓ 1053ms | http |
| 121.130.199.80:3128 | ✓ 1538ms | ✓ 1738ms | ✓ 1750ms | 否 | ✓ 1310ms | http |
| 115.231.181.40:8128 | ✓ 1122ms | ✓ 1050ms | 否 | ✓ 1470ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1989ms | ✓ 1326ms | 否 | ✓ 1336ms | ✓ 1033ms | http |
| 34.71.229.255:3128 | ✓ 501ms | ✓ 1449ms | ✓ 1956ms | ✓ 1424ms | ✓ 1390ms | http |
| 218.108.131.186:17890 | ✓ 1535ms | ✓ 1364ms | ✓ 901ms | ✓ 1092ms | 否 | http |
| 46.101.95.183:8888 | ✓ 1025ms | 否 | ✓ 1470ms | ✓ 1950ms | 否 | http |
| 160.250.4.245:1 | ✓ 1936ms | 否 | ✓ 1395ms | ✓ 1546ms | ✓ 1176ms | http |
| 8.219.195.129:1080 | ✓ 893ms | ✓ 1717ms | ✓ 852ms | ✓ 1050ms | ✓ 839ms | http |
| 113.176.92.71:3128 | ✓ 1530ms | ✓ 1307ms | ✓ 1191ms | ✓ 1189ms | ✓ 921ms | http |
| 118.31.1.154:80 | ✓ 1928ms | ✓ 1490ms | ✓ 855ms | ✓ 1126ms | 否 | http |
| 206.206.126.177:2412 | 否 | 否 | ✓ 744ms | ✓ 1000ms | ✓ 815ms | http |
| 130.61.174.200:1080 | ✓ 1825ms | ✓ 1594ms | 否 | 否 | ✓ 1285ms | http |
| 8.243.162.34:3128 | 否 | ✓ 1695ms | 否 | ✓ 1855ms | ✓ 1849ms | http |
| 38.180.192.119:3128 | ✓ 1813ms | ✓ 798ms | ✓ 762ms | ✓ 845ms | ✓ 579ms | http |
| 161.35.181.96:999 | ✓ 1192ms | ✓ 1452ms | ✓ 588ms | ✓ 1305ms | ✓ 1136ms | http |
| 103.55.225.34:8080 | ✓ 1298ms | 否 | ✓ 1248ms | ✓ 1396ms | ✓ 1354ms | http |
| 110.76.145.116:8090 | ✓ 1788ms | 否 | ✓ 1410ms | ✓ 1304ms | ✓ 1404ms | http |
| 126.209.18.142:8082 | ✓ 1479ms | 否 | ✓ 1997ms | ✓ 1462ms | ✓ 1446ms | http |
| 180.191.233.125:8080 | ✓ 1628ms | 否 | ✓ 1852ms | ✓ 1813ms | ✓ 1591ms | http |
| 121.232.73.214:1080 | 否 | 否 | ✓ 1417ms | ✓ 1484ms | ✓ 1339ms | http |
| 208.87.243.199:7878 | ✓ 1094ms | ✓ 1660ms | ✓ 1367ms | 否 | 否 | http |
| 49.48.41.7:8080 | ✓ 1369ms | 否 | ✓ 1385ms | ✓ 1539ms | 否 | http |
| 177.93.132.244:3128 | ✓ 820ms | 否 | ✓ 893ms | 否 | ✓ 1802ms | http |
| 120.92.212.16:8890 | ✓ 985ms | 否 | ✓ 1934ms | ✓ 1443ms | ✓ 1138ms | http |
| 85.239.59.252:7890 | ✓ 1870ms | ✓ 1832ms | ✓ 1167ms | 否 | 否 | http |
| 36.156.142.27:8083 | ✓ 628ms | ✓ 927ms | ✓ 774ms | ✓ 850ms | ✓ 697ms | http |
| 183.232.248.73:7890 | ✓ 833ms | ✓ 1133ms | ✓ 961ms | ✓ 1073ms | ✓ 846ms | http |
| 120.92.212.16:7890 | ✓ 1987ms | 否 | ✓ 1299ms | ✓ 1279ms | ✓ 1585ms | http |
| 160.238.65.8:3128 | ✓ 876ms | ✓ 1678ms | ✓ 976ms | 否 | ✓ 1236ms | http |
| 160.238.65.5:3128 | 否 | ✓ 1760ms | ✓ 713ms | 否 | ✓ 1662ms | http |
| 160.238.65.3:3128 | 否 | ✓ 1579ms | ✓ 1651ms | ✓ 1600ms | 否 | http |
| 160.238.65.7:3128 | ✓ 673ms | ✓ 1927ms | ✓ 1755ms | 否 | ✓ 1282ms | http |
| 23.224.193.43:3128 | ✓ 503ms | 否 | ✓ 1754ms | ✓ 1254ms | ✓ 1538ms | http |
| 23.224.193.46:3128 | ✓ 1089ms | ✓ 1201ms | ✓ 1671ms | ✓ 1399ms | 否 | http |
| 23.224.193.45:3128 | ✓ 1165ms | ✓ 1177ms | ✓ 1617ms | ✓ 1379ms | ✓ 1909ms | http |
| 23.224.193.42:3128 | ✓ 1403ms | 否 | ✓ 558ms | ✓ 1404ms | ✓ 1869ms | http |
| 106.44.155.90:2222 | ✓ 1516ms | ✓ 1616ms | ✓ 1811ms | 否 | ✓ 1328ms | http |
| 121.230.8.136:1080 | ✓ 1097ms | ✓ 1443ms | ✓ 1242ms | ✓ 1359ms | ✓ 1339ms | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 1365ms | ✓ 811ms | ✓ 1101ms | http |
| 34.96.238.40:8080 | ✓ 815ms | 否 | ✓ 1125ms | 否 | ✓ 1009ms | http |
| 45.140.147.155:1082 | ✓ 1686ms | 否 | ✓ 1974ms | 否 | ✓ 1555ms | http |
| 61.52.131.172:8443 | ✓ 856ms | 否 | 否 | ✓ 1777ms | ✓ 1435ms | http |
| 18.171.232.214:23578 | ✓ 1042ms | 否 | ✓ 1790ms | 否 | ✓ 1900ms | http |
| 13.38.217.179:36632 | ✓ 1003ms | 否 | ✓ 1816ms | 否 | ✓ 1865ms | http |
| 3.121.130.230:56598 | ✓ 1971ms | 否 | ✓ 1015ms | 否 | ✓ 1853ms | http |
| 121.138.61.78:8966 | ✓ 1106ms | ✓ 1697ms | ✓ 1588ms | ✓ 1232ms | ✓ 932ms | http |

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
