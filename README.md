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

最后更新：2026-03-08 06:41:31 UTC（2026-03-08 14:41:31 UTC+8）

**代理总数：70**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1854ms | ✓ 1304ms | ✓ 846ms | ✓ 901ms | ✓ 666ms | http |
| 205.209.118.30:3138 | ✓ 1060ms | 否 | ✓ 944ms | ✓ 1372ms | ✓ 1051ms | http |
| 14.56.107.244:3128 | 否 | 否 | ✓ 907ms | ✓ 1207ms | ✓ 750ms | http |
| 35.225.22.61:80 | ✓ 467ms | ✓ 1511ms | ✓ 791ms | ✓ 1178ms | ✓ 923ms | http |
| 46.183.25.8:443 | ✓ 1777ms | 否 | ✓ 1095ms | ✓ 1614ms | 否 | http |
| 202.155.12.161:443 | ✓ 1850ms | 否 | ✓ 1198ms | ✓ 1221ms | ✓ 953ms | http |
| 152.42.213.210:8080 | ✓ 1421ms | 否 | ✓ 777ms | ✓ 1453ms | 否 | http |
| 62.113.119.14:8080 | ✓ 947ms | 否 | ✓ 1427ms | ✓ 1702ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1451ms | ✓ 943ms | 否 | 否 | ✓ 841ms | http |
| 120.92.212.16:7890 | ✓ 990ms | ✓ 1723ms | ✓ 794ms | 否 | 否 | http |
| 159.223.42.219:3128 | ✓ 1692ms | 否 | 否 | ✓ 1619ms | ✓ 979ms | http |
| 168.235.110.63:3128 | 否 | ✓ 1272ms | ✓ 878ms | ✓ 1308ms | ✓ 1004ms | http |
| 81.70.169.194:80 | 否 | ✓ 1520ms | 否 | ✓ 1154ms | ✓ 740ms | http |
| 101.43.255.96:80 | ✓ 893ms | 否 | ✓ 735ms | ✓ 1246ms | ✓ 715ms | http |
| 59.46.216.131:30001 | ✓ 862ms | 否 | ✓ 1109ms | 否 | ✓ 1112ms | http |
| 159.89.31.62:8080 | ✓ 1574ms | ✓ 1838ms | ✓ 1975ms | 否 | 否 | http |
| 165.227.5.10:8888 | 否 | ✓ 1286ms | ✓ 1379ms | ✓ 787ms | 否 | http |
| 103.82.23.118:5247 | ✓ 1970ms | 否 | ✓ 1806ms | 否 | ✓ 1862ms | http |
| 116.80.82.218:3172 | ✓ 1452ms | 否 | ✓ 1510ms | ✓ 1808ms | 否 | http |
| 116.80.82.231:3172 | ✓ 1453ms | 否 | ✓ 1509ms | 否 | ✓ 1619ms | http |
| 47.101.159.19:8899 | ✓ 761ms | ✓ 777ms | ✓ 754ms | ✓ 966ms | ✓ 717ms | http |
| 39.104.201.40:7890 | ✓ 1058ms | ✓ 1295ms | ✓ 1349ms | ✓ 1006ms | ✓ 1652ms | http |
| 120.92.212.16:8890 | ✓ 1789ms | 否 | 否 | ✓ 1193ms | ✓ 1206ms | http |
| 46.249.103.192:443 | ✓ 1286ms | 否 | ✓ 1661ms | ✓ 1639ms | 否 | http |
| 61.72.110.54:3128 | ✓ 622ms | 否 | ✓ 1102ms | 否 | ✓ 1511ms | http |
| 200.174.198.32:8888 | ✓ 1544ms | 否 | ✓ 1660ms | 否 | ✓ 1868ms | http |
| 116.80.82.230:3172 | ✓ 1463ms | 否 | ✓ 1473ms | ✓ 1760ms | 否 | http |
| 103.84.95.54:7890 | ✓ 755ms | 否 | ✓ 1130ms | 否 | ✓ 607ms | http |
| 14.225.222.185:7890 | ✓ 955ms | 否 | ✓ 832ms | ✓ 1103ms | 否 | http |
| 125.128.12.14:3128 | 否 | 否 | ✓ 1769ms | ✓ 1012ms | ✓ 1320ms | http |
| 103.139.138.194:3128 | ✓ 1796ms | 否 | ✓ 1722ms | ✓ 1473ms | ✓ 1422ms | http |
| 67.169.98.211:443 | ✓ 1442ms | 否 | ✓ 1057ms | ✓ 1338ms | ✓ 1433ms | http |
| 61.72.221.194:3128 | 否 | ✓ 1706ms | ✓ 1497ms | ✓ 1424ms | ✓ 1757ms | http |
| 116.80.82.224:3172 | ✓ 1558ms | 否 | 否 | ✓ 1825ms | ✓ 1599ms | http |
| 178.236.245.17:3128 | ✓ 782ms | 否 | ✓ 1419ms | 否 | ✓ 1630ms | http |
| 178.236.245.59:3128 | ✓ 790ms | ✓ 1999ms | ✓ 1425ms | 否 | ✓ 1579ms | http |
| 121.204.158.249:3128 | 否 | ✓ 1585ms | 否 | ✓ 1418ms | ✓ 1463ms | http |
| 185.243.218.43:49153 | ✓ 1301ms | 否 | ✓ 1606ms | 否 | ✓ 1632ms | http |
| 103.82.23.118:5234 | ✓ 1535ms | 否 | ✓ 1601ms | ✓ 1752ms | ✓ 1802ms | http |
| 121.230.9.248:1080 | 否 | ✓ 1756ms | ✓ 1290ms | 否 | ✓ 1426ms | http |
| 103.82.23.118:5171 | ✓ 1702ms | 否 | ✓ 1775ms | ✓ 1813ms | ✓ 1583ms | http |
| 61.109.216.213:8080 | ✓ 1115ms | ✓ 1035ms | ✓ 1582ms | ✓ 1529ms | ✓ 928ms | http |
| 45.136.198.40:3128 | ✓ 867ms | ✓ 1708ms | ✓ 845ms | ✓ 1692ms | ✓ 1394ms | http |
| 16.78.119.130:443 | ✓ 1975ms | ✓ 1936ms | 否 | ✓ 1749ms | 否 | http |
| 121.128.121.54:3128 | 否 | 否 | ✓ 591ms | ✓ 956ms | ✓ 1105ms | http |
| 150.107.140.238:3128 | ✓ 1481ms | 否 | 否 | ✓ 1242ms | ✓ 886ms | http |
| 85.9.195.140:1080 | ✓ 1772ms | 否 | 否 | ✓ 1528ms | ✓ 1557ms | http |
| 1.225.116.115:1080 | ✓ 695ms | 否 | ✓ 896ms | ✓ 966ms | ✓ 774ms | http |
| 138.124.53.25:7443 | ✓ 1303ms | 否 | 否 | ✓ 1862ms | ✓ 1495ms | http |
| 163.44.126.97:3128 | ✓ 980ms | 否 | ✓ 1362ms | ✓ 1280ms | ✓ 1019ms | http |
| 103.183.10.169:3125 | ✓ 1810ms | 否 | ✓ 1886ms | ✓ 1876ms | 否 | http |
| 103.215.36.88:17148 | ✓ 908ms | ✓ 1174ms | ✓ 1164ms | ✓ 1373ms | ✓ 906ms | http |
| 103.215.36.88:18112 | ✓ 894ms | 否 | ✓ 1051ms | ✓ 1269ms | ✓ 1593ms | http |
| 34.96.238.40:8080 | ✓ 1479ms | ✓ 1636ms | ✓ 1039ms | ✓ 970ms | ✓ 801ms | http |
| 61.72.221.94:3128 | ✓ 1506ms | ✓ 1818ms | ✓ 1623ms | ✓ 1381ms | ✓ 1320ms | http |
| 103.39.51.190:8080 | ✓ 1959ms | 否 | 否 | ✓ 1986ms | ✓ 1400ms | http |
| 103.215.36.88:18821 | ✓ 1345ms | ✓ 1646ms | ✓ 1155ms | 否 | ✓ 976ms | http |
| 103.215.36.88:19385 | ✓ 1021ms | ✓ 1159ms | 否 | ✓ 1285ms | 否 | http |
| 103.236.89.228:7890 | ✓ 844ms | ✓ 1584ms | ✓ 1534ms | 否 | ✓ 992ms | http |
| 193.228.139.78:8888 | ✓ 895ms | 否 | ✓ 1856ms | 否 | ✓ 1818ms | http |
| 116.80.82.226:3172 | ✓ 1477ms | 否 | ✓ 1502ms | 否 | ✓ 1632ms | http |
| 218.60.0.214:80 | ✓ 1472ms | ✓ 1700ms | ✓ 1194ms | ✓ 1483ms | 否 | http |
| 113.177.131.2:3128 | ✓ 1018ms | 否 | ✓ 1088ms | ✓ 1643ms | ✓ 875ms | http |
| 82.65.27.56:80 | ✓ 1411ms | 否 | ✓ 1201ms | ✓ 1642ms | 否 | http |
| 45.140.147.82:1081 | ✓ 1254ms | 否 | ✓ 1368ms | 否 | ✓ 1193ms | http |
| 112.78.187.186:8080 | ✓ 1203ms | 否 | ✓ 1428ms | ✓ 1347ms | ✓ 1311ms | http |
| 114.4.251.26:8080 | ✓ 1213ms | 否 | ✓ 1251ms | ✓ 1387ms | ✓ 1496ms | http |
| 45.140.147.82:1082 | ✓ 760ms | ✓ 1612ms | ✓ 1110ms | 否 | ✓ 1408ms | http |
| 106.14.203.63:3333 | ✓ 1527ms | 否 | 否 | ✓ 820ms | ✓ 1413ms | http |
| 54.222.174.194:80 | 否 | ✓ 1375ms | ✓ 1362ms | 否 | ✓ 1176ms | http |

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
