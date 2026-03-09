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

最后更新：2026-03-09 11:54:49 UTC（2026-03-09 19:54:49 UTC+8）

**代理总数：54**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 53 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 54 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | 否 | ✓ 1512ms | ✓ 1820ms | ✓ 1550ms | ✓ 1241ms | http |
| 45.140.147.82:1081 | ✓ 883ms | ✓ 1614ms | ✓ 850ms | ✓ 1915ms | ✓ 1303ms | http |
| 101.47.73.135:3128 | ✓ 1565ms | 否 | ✓ 1145ms | ✓ 1500ms | ✓ 1334ms | http |
| 35.225.22.61:80 | ✓ 876ms | ✓ 1138ms | ✓ 394ms | ✓ 1075ms | ✓ 680ms | http |
| 64.186.232.4:10808 | 否 | 否 | ✓ 961ms | ✓ 981ms | ✓ 1755ms | http |
| 194.213.18.200:443 | ✓ 1670ms | 否 | 否 | ✓ 1636ms | ✓ 1663ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 1984ms | ✓ 1132ms | ✓ 807ms | http |
| 162.248.165.72:1080 | ✓ 1525ms | 否 | ✓ 891ms | 否 | ✓ 1685ms | http |
| 61.72.221.94:3128 | ✓ 1295ms | 否 | ✓ 1313ms | ✓ 1124ms | ✓ 939ms | http |
| 121.237.181.137:8888 | ✓ 941ms | ✓ 1301ms | ✓ 973ms | ✓ 1362ms | ✓ 1894ms | http |
| 168.235.110.63:3128 | ✓ 691ms | 否 | 否 | ✓ 1478ms | ✓ 936ms | http |
| 81.70.169.194:80 | 否 | 否 | ✓ 1284ms | ✓ 1378ms | ✓ 1735ms | http |
| 120.92.212.16:8890 | ✓ 1862ms | 否 | ✓ 1326ms | ✓ 1403ms | ✓ 1329ms | http |
| 101.43.255.96:80 | ✓ 1114ms | ✓ 1506ms | ✓ 1021ms | 否 | 否 | http |
| 107.172.125.217:3128 | ✓ 274ms | 否 | ✓ 713ms | ✓ 884ms | ✓ 735ms | http |
| 14.56.177.44:3128 | ✓ 1269ms | ✓ 1433ms | ✓ 1305ms | ✓ 1262ms | ✓ 1004ms | http |
| 178.236.245.17:3128 | ✓ 1014ms | 否 | ✓ 1638ms | 否 | ✓ 1676ms | http |
| 202.155.12.161:443 | ✓ 1741ms | 否 | ✓ 1303ms | ✓ 1132ms | 否 | http |
| 173.212.246.157:3128 | ✓ 1352ms | 否 | ✓ 710ms | ✓ 1861ms | ✓ 1609ms | http |
| 120.92.212.16:7890 | ✓ 1343ms | 否 | ✓ 1347ms | 否 | ✓ 1049ms | http |
| 159.223.42.219:3128 | ✓ 880ms | 否 | ✓ 878ms | ✓ 1329ms | 否 | http |
| 38.180.2.107:3128 | ✓ 1377ms | 否 | ✓ 1761ms | ✓ 1897ms | ✓ 1712ms | http |
| 159.89.31.62:8080 | ✓ 1215ms | ✓ 1758ms | ✓ 1562ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 1295ms | ✓ 1881ms | ✓ 1872ms | ✓ 1987ms | ✓ 1634ms | http |
| 62.113.119.14:8080 | ✓ 872ms | 否 | 否 | ✓ 1865ms | ✓ 1263ms | http |
| 14.225.212.37:7890 | ✓ 1509ms | 否 | ✓ 1147ms | ✓ 1186ms | ✓ 1615ms | http |
| 205.209.118.30:3138 | 否 | 否 | ✓ 1881ms | ✓ 1720ms | ✓ 1890ms | http |
| 202.129.206.239:3128 | ✓ 1457ms | 否 | 否 | ✓ 1854ms | ✓ 1662ms | http |
| 209.38.51.97:3128 | ✓ 300ms | 否 | ✓ 942ms | 否 | ✓ 823ms | http |
| 178.236.245.59:3128 | ✓ 1180ms | 否 | ✓ 1471ms | ✓ 1987ms | ✓ 1299ms | http |
| 58.220.95.11:11023 | ✓ 1015ms | ✓ 1340ms | ✓ 1022ms | ✓ 1427ms | ✓ 1297ms | http |
| 136.49.34.18:8888 | ✓ 1338ms | ✓ 1926ms | 否 | ✓ 1395ms | ✓ 1755ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1235ms | ✓ 1036ms | ✓ 1426ms | ✓ 1022ms | http |
| 172.212.68.37:3128 | ✓ 365ms | ✓ 1441ms | ✓ 738ms | ✓ 1554ms | ✓ 1091ms | http |
| 190.9.109.198:999 | ✓ 697ms | 否 | ✓ 1490ms | ✓ 1445ms | 否 | http |
| 211.171.114.154:3128 | ✓ 1763ms | ✓ 1313ms | 否 | ✓ 1581ms | ✓ 1242ms | http |
| 152.42.213.210:8080 | ✓ 878ms | 否 | ✓ 1456ms | ✓ 1677ms | ✓ 1347ms | http |
| 54.222.174.194:80 | 否 | ✓ 1762ms | ✓ 1807ms | 否 | ✓ 1717ms | http |
| 91.233.223.147:3128 | ✓ 1210ms | 否 | ✓ 1610ms | ✓ 1946ms | 否 | http |
| 120.55.163.237:10086 | ✓ 1032ms | ✓ 1235ms | ✓ 1088ms | ✓ 1259ms | ✓ 1009ms | http |
| 47.95.231.180:8084 | ✓ 983ms | ✓ 1313ms | ✓ 988ms | ✓ 1331ms | ✓ 1048ms | http |
| 116.80.96.104:3172 | ✓ 1585ms | 否 | ✓ 1269ms | ✓ 1139ms | ✓ 966ms | http |
| 88.80.150.82:8080 | ✓ 1702ms | ✓ 1916ms | 否 | 否 | ✓ 1742ms | https |
| 39.104.201.40:7890 | 否 | 否 | ✓ 1015ms | ✓ 1333ms | ✓ 1077ms | http |
| 103.39.51.190:8080 | ✓ 1845ms | 否 | 否 | ✓ 1756ms | ✓ 1626ms | http |
| 34.96.238.40:8080 | ✓ 1309ms | ✓ 1239ms | ✓ 1270ms | 否 | ✓ 952ms | http |
| 192.71.213.85:9091 | ✓ 731ms | 否 | ✓ 875ms | ✓ 1660ms | 否 | http |
| 1.225.116.115:1080 | ✓ 1659ms | ✓ 1598ms | 否 | 否 | ✓ 1040ms | http |
| 213.220.3.183:3128 | ✓ 1416ms | ✓ 1728ms | ✓ 1724ms | 否 | 否 | http |
| 116.6.106.33:3128 | ✓ 1130ms | ✓ 1270ms | ✓ 935ms | 否 | 否 | http |
| 201.144.20.238:3128 | 否 | ✓ 1996ms | 否 | ✓ 1229ms | ✓ 961ms | http |
| 34.101.184.164:3128 | ✓ 1716ms | 否 | ✓ 1603ms | ✓ 1717ms | ✓ 1121ms | http |
| 103.178.86.178:8080 | ✓ 1573ms | 否 | ✓ 1603ms | ✓ 1750ms | ✓ 1644ms | http |
| 150.230.211.74:8080 | ✓ 1626ms | 否 | ✓ 1546ms | 否 | ✓ 1323ms | http |

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
