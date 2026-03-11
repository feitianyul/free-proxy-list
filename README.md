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

最后更新：2026-03-11 14:05:24 UTC（2026-03-11 22:05:24 UTC+8）

**代理总数：64**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 432ms | ✓ 1264ms | 否 | ✓ 1316ms | 否 | http |
| 45.136.131.63:8443 | ✓ 365ms | ✓ 899ms | ✓ 1109ms | ✓ 827ms | ✓ 577ms | http |
| 1.231.81.166:3128 | ✓ 1168ms | 否 | ✓ 891ms | ✓ 802ms | ✓ 657ms | http |
| 202.155.12.161:443 | ✓ 1760ms | 否 | 否 | ✓ 1942ms | ✓ 1641ms | http |
| 45.136.198.40:3128 | ✓ 926ms | 否 | ✓ 1914ms | 否 | ✓ 1941ms | http |
| 165.227.5.10:8888 | ✓ 1619ms | 否 | 否 | ✓ 902ms | ✓ 603ms | http |
| 45.136.131.47:8443 | ✓ 281ms | ✓ 707ms | ✓ 954ms | ✓ 789ms | ✓ 727ms | http |
| 45.136.130.223:8443 | ✓ 187ms | ✓ 1393ms | ✓ 774ms | ✓ 895ms | ✓ 617ms | http |
| 45.136.130.175:8443 | ✓ 337ms | ✓ 1932ms | ✓ 780ms | ✓ 883ms | ✓ 828ms | http |
| 158.69.185.37:3129 | ✓ 631ms | 否 | ✓ 945ms | ✓ 1455ms | ✓ 1042ms | http |
| 162.240.154.26:3128 | ✓ 794ms | 否 | ✓ 1286ms | ✓ 1718ms | ✓ 1319ms | http |
| 167.172.77.49:8080 | ✓ 1987ms | 否 | ✓ 842ms | ✓ 1836ms | ✓ 1120ms | http |
| 115.231.181.40:8128 | ✓ 661ms | 否 | ✓ 729ms | 否 | ✓ 1764ms | http |
| 120.92.212.16:8890 | ✓ 896ms | ✓ 1179ms | 否 | ✓ 993ms | ✓ 804ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1885ms | ✓ 1188ms | ✓ 784ms | http |
| 45.136.130.191:8443 | ✓ 166ms | ✓ 747ms | ✓ 378ms | ✓ 781ms | ✓ 610ms | http |
| 45.136.130.239:8443 | 否 | ✓ 792ms | 否 | ✓ 1131ms | ✓ 589ms | http |
| 152.70.98.46:8888 | ✓ 1112ms | 否 | ✓ 1494ms | ✓ 941ms | ✓ 1484ms | http |
| 101.43.255.96:80 | ✓ 1418ms | ✓ 1112ms | ✓ 1083ms | 否 | ✓ 896ms | http |
| 46.183.25.8:443 | ✓ 657ms | 否 | ✓ 453ms | ✓ 1378ms | 否 | http |
| 194.213.18.200:443 | ✓ 1091ms | 否 | ✓ 1143ms | ✓ 1301ms | 否 | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1201ms | ✓ 1460ms | ✓ 1613ms | http |
| 91.107.141.42:8081 | ✓ 677ms | 否 | ✓ 1265ms | 否 | ✓ 1759ms | http |
| 162.248.165.72:1080 | ✓ 1479ms | 否 | ✓ 1837ms | 否 | ✓ 1573ms | http |
| 39.104.201.40:7890 | ✓ 660ms | ✓ 927ms | 否 | ✓ 925ms | ✓ 736ms | http |
| 81.70.169.194:80 | ✓ 1805ms | 否 | ✓ 1317ms | ✓ 982ms | ✓ 1090ms | http |
| 113.177.131.2:3128 | ✓ 1853ms | 否 | ✓ 1094ms | ✓ 1388ms | ✓ 979ms | http |
| 210.223.44.230:3128 | ✓ 1475ms | ✓ 1740ms | ✓ 1489ms | ✓ 1938ms | 否 | http |
| 122.248.45.54:8080 | ✓ 1819ms | 否 | ✓ 1562ms | ✓ 1499ms | 否 | http |
| 116.80.82.224:3172 | ✓ 1753ms | 否 | 否 | ✓ 1863ms | ✓ 1582ms | http |
| 47.101.159.19:8899 | ✓ 682ms | ✓ 737ms | ✓ 608ms | ✓ 797ms | ✓ 727ms | http |
| 94.176.3.43:7443 | ✓ 1896ms | 否 | ✓ 1536ms | 否 | ✓ 1557ms | http |
| 113.176.92.71:3128 | ✓ 1964ms | ✓ 1234ms | ✓ 1298ms | ✓ 1086ms | ✓ 880ms | http |
| 178.236.245.17:3128 | ✓ 1058ms | ✓ 1845ms | ✓ 1188ms | ✓ 1971ms | 否 | http |
| 190.9.109.198:999 | ✓ 923ms | 否 | ✓ 1460ms | ✓ 1647ms | 否 | http |
| 185.191.236.162:3128 | ✓ 1388ms | 否 | ✓ 1827ms | 否 | ✓ 1789ms | http |
| 111.48.191.1:7890 | ✓ 676ms | ✓ 955ms | ✓ 681ms | ✓ 965ms | ✓ 781ms | http |
| 103.82.93.219:3128 | ✓ 1614ms | 否 | ✓ 1804ms | ✓ 1152ms | ✓ 914ms | http |
| 121.230.8.34:1080 | ✓ 1008ms | ✓ 1249ms | ✓ 962ms | 否 | 否 | http |
| 54.222.174.194:80 | ✓ 1662ms | 否 | ✓ 1479ms | ✓ 1352ms | 否 | http |
| 47.101.149.27:9010 | ✓ 1011ms | ✓ 857ms | 否 | ✓ 982ms | 否 | http |
| 42.96.16.158:1311 | 否 | 否 | ✓ 1997ms | ✓ 1194ms | ✓ 1012ms | http |
| 62.113.119.14:8080 | ✓ 1279ms | 否 | ✓ 1428ms | ✓ 1866ms | ✓ 1495ms | http |
| 59.46.216.131:30001 | ✓ 827ms | ✓ 1128ms | 否 | 否 | ✓ 862ms | http |
| 14.225.222.213:7890 | ✓ 1093ms | 否 | ✓ 1068ms | ✓ 1263ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1187ms | ✓ 1992ms | ✓ 1888ms | 否 | 否 | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1504ms | ✓ 1616ms | ✓ 1224ms | http |
| 45.136.130.188:8443 | ✓ 508ms | 否 | ✓ 1533ms | ✓ 1374ms | ✓ 647ms | http |
| 103.82.23.118:5221 | ✓ 1767ms | 否 | ✓ 1535ms | ✓ 1957ms | ✓ 1255ms | http |
| 34.101.184.164:3128 | ✓ 1646ms | 否 | ✓ 1240ms | ✓ 1206ms | ✓ 1076ms | http |
| 103.39.51.190:8080 | ✓ 1621ms | 否 | 否 | ✓ 1321ms | ✓ 1283ms | http |
| 95.3.9.78:3128 | ✓ 1496ms | 否 | ✓ 1991ms | 否 | ✓ 1458ms | http |
| 61.52.131.172:8443 | ✓ 630ms | ✓ 858ms | ✓ 824ms | ✓ 862ms | ✓ 656ms | http |
| 190.6.54.12:6969 | ✓ 1182ms | 否 | 否 | ✓ 1937ms | ✓ 1878ms | http |
| 95.3.9.78:8080 | 否 | 否 | ✓ 1290ms | ✓ 1870ms | ✓ 1569ms | http |
| 116.80.96.105:3172 | ✓ 404ms | 否 | ✓ 449ms | ✓ 714ms | ✓ 562ms | http |
| 116.80.96.106:3172 | ✓ 401ms | 否 | ✓ 452ms | ✓ 728ms | ✓ 591ms | http |
| 116.80.96.111:3172 | ✓ 425ms | 否 | ✓ 429ms | ✓ 815ms | ✓ 661ms | http |
| 107.173.52.58:7890 | ✓ 382ms | ✓ 1941ms | ✓ 1128ms | ✓ 1334ms | ✓ 1066ms | http |
| 149.88.94.216:7890 | ✓ 751ms | 否 | ✓ 657ms | ✓ 869ms | ✓ 651ms | http |
| 113.160.132.26:8080 | ✓ 876ms | ✓ 1125ms | ✓ 1189ms | ✓ 1073ms | ✓ 883ms | http |
| 116.80.96.110:3172 | ✓ 1406ms | 否 | ✓ 1460ms | 否 | ✓ 1625ms | http |
| 8.219.97.248:80 | ✓ 1355ms | 否 | ✓ 1849ms | ✓ 1515ms | 否 | http |
| 106.14.203.63:3333 | 否 | ✓ 1030ms | 否 | ✓ 1904ms | ✓ 902ms | http |

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
