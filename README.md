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

最后更新：2026-04-24 07:02:01 UTC（2026-04-24 15:02:01 UTC+8）

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
| 47.85.51.197:1080 | ✓ 544ms | ✓ 1065ms | ✓ 706ms | ✓ 897ms | ✓ 984ms | http |
| 1.231.81.166:3128 | ✓ 1787ms | ✓ 1284ms | ✓ 1113ms | ✓ 1149ms | ✓ 1033ms | http |
| 46.101.95.183:8888 | 否 | 否 | ✓ 1361ms | ✓ 1437ms | ✓ 1169ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1668ms | ✓ 1548ms | ✓ 1709ms | 否 | http |
| 35.225.22.61:80 | ✓ 319ms | ✓ 1939ms | ✓ 1095ms | 否 | 否 | http |
| 8.209.238.110:47701 | ✓ 828ms | 否 | ✓ 812ms | ✓ 1034ms | ✓ 884ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1134ms | ✓ 1501ms | ✓ 1200ms | http |
| 45.153.231.229:8080 | ✓ 1074ms | 否 | ✓ 1917ms | 否 | ✓ 1803ms | http |
| 20.127.128.70:8080 | ✓ 1755ms | 否 | ✓ 1101ms | 否 | ✓ 1915ms | http |
| 45.140.147.155:1082 | ✓ 845ms | 否 | ✓ 1038ms | 否 | ✓ 1280ms | http |
| 89.208.106.138:10808 | 否 | ✓ 1542ms | 否 | ✓ 1489ms | ✓ 1304ms | http |
| 146.19.56.212:40002 | 否 | 否 | ✓ 343ms | ✓ 1659ms | ✓ 1410ms | http |
| 130.61.139.145:3128 | ✓ 693ms | 否 | ✓ 1820ms | ✓ 1688ms | ✓ 1138ms | http |
| 8.219.195.129:1080 | ✓ 1213ms | ✓ 1946ms | ✓ 897ms | ✓ 1304ms | ✓ 1048ms | http |
| 47.95.231.180:8084 | ✓ 1063ms | ✓ 1438ms | ✓ 1114ms | ✓ 1383ms | ✓ 1129ms | http |
| 91.99.15.45:2095 | ✓ 1451ms | ✓ 1926ms | 否 | 否 | ✓ 1495ms | http |
| 47.110.42.192:9003 | ✓ 1639ms | ✓ 1627ms | ✓ 1775ms | ✓ 1826ms | ✓ 1573ms | http |
| 168.144.75.9:3128 | ✓ 1655ms | 否 | 否 | ✓ 1918ms | ✓ 1449ms | http |
| 59.46.216.131:30001 | ✓ 1269ms | 否 | ✓ 1942ms | 否 | ✓ 1326ms | http |
| 45.140.147.82:1082 | ✓ 527ms | ✓ 1172ms | ✓ 696ms | ✓ 1597ms | ✓ 1155ms | http |
| 45.140.147.82:1081 | ✓ 536ms | ✓ 1504ms | ✓ 460ms | ✓ 1471ms | ✓ 1186ms | http |
| 20.164.75.153:8080 | ✓ 1151ms | 否 | ✓ 1183ms | 否 | ✓ 1998ms | http |
| 34.96.238.40:8080 | ✓ 1268ms | 否 | ✓ 1228ms | ✓ 1232ms | 否 | http |
| 193.181.35.169:8118 | 否 | 否 | ✓ 1079ms | ✓ 1922ms | ✓ 1483ms | http |
| 161.35.181.96:999 | ✓ 522ms | ✓ 962ms | ✓ 308ms | ✓ 1561ms | ✓ 1102ms | http |
| 160.19.19.139:8097 | 否 | 否 | ✓ 1797ms | ✓ 1737ms | ✓ 1708ms | http |
| 135.125.97.184:40551 | ✓ 997ms | ✓ 1774ms | 否 | 否 | ✓ 1817ms | http |
| 159.89.31.62:8080 | ✓ 857ms | 否 | ✓ 1453ms | ✓ 1866ms | ✓ 1275ms | http |
| 152.42.177.32:8888 | ✓ 1191ms | 否 | ✓ 1121ms | ✓ 1520ms | ✓ 1534ms | http |
| 130.61.174.200:1080 | 否 | ✓ 1281ms | ✓ 539ms | 否 | ✓ 1918ms | http |
| 85.190.99.143:443 | ✓ 1178ms | 否 | ✓ 1937ms | 否 | ✓ 1690ms | http |
| 20.78.213.56:80 | ✓ 660ms | ✓ 1724ms | ✓ 938ms | ✓ 1149ms | 否 | http |
| 94.241.173.165:1080 | 否 | 否 | ✓ 1706ms | ✓ 1899ms | ✓ 1262ms | http |
| 177.93.132.244:3128 | ✓ 780ms | 否 | ✓ 1908ms | 否 | ✓ 1684ms | http |
| 115.231.181.40:8128 | ✓ 1212ms | ✓ 1301ms | 否 | 否 | ✓ 1219ms | http |
| 144.31.25.69:21064 | ✓ 1383ms | 否 | ✓ 1347ms | 否 | ✓ 1960ms | http |
| 47.84.73.61:1080 | ✓ 1458ms | 否 | ✓ 955ms | ✓ 1354ms | ✓ 1038ms | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 1873ms | ✓ 1343ms | ✓ 1329ms | http |
| 2.27.40.180:1080 | ✓ 1977ms | ✓ 1841ms | ✓ 1180ms | 否 | 否 | http |
| 84.47.150.125:1080 | ✓ 931ms | 否 | 否 | ✓ 1502ms | ✓ 1360ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1376ms | ✓ 1654ms | ✓ 1652ms | http |
| 45.140.147.155:1081 | ✓ 1454ms | 否 | ✓ 509ms | ✓ 1265ms | ✓ 918ms | http |
| 152.70.91.193:40000 | ✓ 1756ms | 否 | ✓ 1913ms | ✓ 1452ms | ✓ 1546ms | http |
| 218.108.131.186:17890 | ✓ 995ms | ✓ 1275ms | ✓ 1028ms | ✓ 1360ms | ✓ 1047ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1875ms | ✓ 1946ms | ✓ 1542ms | http |
| 91.193.240.157:9877 | ✓ 1525ms | 否 | ✓ 1044ms | 否 | ✓ 1630ms | http |
| 114.237.77.250:1080 | ✓ 1073ms | ✓ 1489ms | ✓ 1786ms | ✓ 1474ms | ✓ 1268ms | http |
| 114.237.77.239:1080 | ✓ 1152ms | ✓ 1734ms | 否 | ✓ 1477ms | 否 | http |
| 51.79.71.106:8080 | ✓ 847ms | ✓ 1865ms | 否 | ✓ 1925ms | ✓ 1188ms | http |
| 201.144.20.238:3128 | ✓ 1677ms | ✓ 1520ms | ✓ 1375ms | ✓ 1240ms | ✓ 1252ms | http |
| 168.110.52.228:3128 | ✓ 766ms | ✓ 1879ms | ✓ 1397ms | ✓ 1646ms | ✓ 1029ms | http |
| 62.113.119.14:8080 | ✓ 690ms | ✓ 1594ms | ✓ 1172ms | ✓ 1636ms | ✓ 1446ms | http |
| 45.186.6.104:3128 | ✓ 1377ms | ✓ 1767ms | ✓ 1628ms | 否 | 否 | http |
| 152.32.132.190:7890 | ✓ 1392ms | 否 | 否 | ✓ 1593ms | ✓ 1111ms | http |
| 146.59.16.105:3128 | ✓ 1407ms | ✓ 1696ms | ✓ 1960ms | 否 | 否 | http |

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
