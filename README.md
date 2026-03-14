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

最后更新：2026-03-14 10:29:55 UTC（2026-03-14 18:29:55 UTC+8）

**代理总数：58**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 57 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 58 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 216.180.127.45:1080 | ✓ 1035ms | 否 | ✓ 1145ms | ✓ 1340ms | ✓ 1084ms | http |
| 205.209.118.30:3138 | ✓ 739ms | 否 | ✓ 1012ms | ✓ 1080ms | ✓ 888ms | http |
| 202.155.12.161:443 | ✓ 1167ms | 否 | ✓ 1133ms | ✓ 1296ms | 否 | http |
| 86.53.183.16:1080 | ✓ 1191ms | 否 | 否 | ✓ 1893ms | ✓ 1331ms | http |
| 147.45.251.242:8888 | ✓ 1199ms | 否 | ✓ 1541ms | 否 | ✓ 1899ms | http |
| 113.160.132.26:8080 | ✓ 1924ms | ✓ 1532ms | ✓ 1403ms | ✓ 1429ms | ✓ 1136ms | http |
| 5.129.206.247:8888 | ✓ 1126ms | ✓ 1857ms | ✓ 1753ms | 否 | 否 | http |
| 45.88.0.116:3128 | ✓ 1869ms | 否 | ✓ 1755ms | 否 | ✓ 1835ms | http |
| 43.167.227.161:1080 | ✓ 1412ms | 否 | ✓ 1894ms | ✓ 1044ms | ✓ 1020ms | http |
| 210.77.29.245:7890 | ✓ 1028ms | ✓ 1433ms | ✓ 1210ms | ✓ 1632ms | ✓ 1245ms | http |
| 45.167.124.52:8080 | ✓ 812ms | ✓ 1752ms | 否 | ✓ 1500ms | ✓ 1234ms | http |
| 103.84.95.54:7890 | ✓ 854ms | 否 | ✓ 901ms | ✓ 1202ms | ✓ 853ms | http |
| 101.47.73.135:3128 | ✓ 1362ms | 否 | ✓ 1136ms | ✓ 1311ms | ✓ 1395ms | http |
| 186.148.180.46:999 | ✓ 727ms | ✓ 1763ms | ✓ 1297ms | ✓ 1531ms | ✓ 1262ms | http |
| 138.124.53.25:7443 | ✓ 460ms | 否 | ✓ 1917ms | ✓ 1823ms | 否 | http |
| 45.168.238.193:8443 | 否 | 否 | ✓ 1006ms | ✓ 1251ms | ✓ 883ms | http |
| 120.92.212.16:7890 | ✓ 1196ms | ✓ 1488ms | 否 | 否 | ✓ 1175ms | http |
| 35.225.22.61:80 | ✓ 1072ms | ✓ 1311ms | ✓ 1041ms | 否 | 否 | http |
| 45.136.131.42:8447 | ✓ 784ms | ✓ 1504ms | ✓ 460ms | ✓ 1072ms | ✓ 1463ms | http |
| 150.230.249.50:1080 | ✓ 1609ms | 否 | ✓ 1117ms | 否 | ✓ 943ms | http |
| 128.199.120.45:9090 | ✓ 1887ms | 否 | ✓ 1184ms | ✓ 1995ms | ✓ 1737ms | http |
| 81.70.169.194:80 | ✓ 1227ms | 否 | ✓ 1296ms | ✓ 1887ms | 否 | http |
| 24.144.86.173:1080 | ✓ 835ms | ✓ 1042ms | ✓ 1490ms | ✓ 883ms | ✓ 674ms | http |
| 45.136.131.39:8443 | 否 | ✓ 1331ms | ✓ 492ms | ✓ 924ms | ✓ 690ms | http |
| 85.198.96.242:3128 | 否 | 否 | ✓ 1902ms | ✓ 1603ms | ✓ 1186ms | http |
| 91.233.223.147:3128 | ✓ 835ms | 否 | ✓ 748ms | ✓ 1826ms | ✓ 1448ms | http |
| 162.243.149.86:31028 | ✓ 647ms | ✓ 1932ms | 否 | 否 | ✓ 1010ms | http |
| 157.100.54.4:80 | ✓ 869ms | 否 | ✓ 1342ms | ✓ 1683ms | ✓ 1399ms | http |
| 91.247.126.241:2080 | ✓ 883ms | 否 | ✓ 1679ms | 否 | ✓ 1723ms | http |
| 34.101.184.164:3128 | ✓ 1126ms | 否 | ✓ 1666ms | ✓ 1543ms | ✓ 1310ms | http |
| 213.220.62.62:3128 | ✓ 1143ms | ✓ 1495ms | ✓ 1875ms | 否 | 否 | http |
| 45.88.0.99:3128 | ✓ 818ms | 否 | ✓ 832ms | ✓ 1964ms | ✓ 1134ms | http |
| 116.80.96.108:3172 | ✓ 1743ms | 否 | 否 | ✓ 1999ms | ✓ 1866ms | http |
| 116.80.49.169:3172 | ✓ 1723ms | 否 | ✓ 1689ms | 否 | ✓ 1877ms | http |
| 101.43.255.96:80 | ✓ 1318ms | 否 | ✓ 1586ms | ✓ 1890ms | 否 | http |
| 45.88.0.114:3128 | ✓ 998ms | 否 | ✓ 722ms | ✓ 1355ms | ✓ 1040ms | http |
| 45.88.0.117:3128 | ✓ 1077ms | 否 | ✓ 643ms | ✓ 1338ms | 否 | http |
| 88.80.150.82:8080 | ✓ 1207ms | 否 | 否 | ✓ 1853ms | ✓ 1706ms | https |
| 162.240.154.26:3128 | ✓ 1049ms | 否 | ✓ 1019ms | ✓ 1479ms | ✓ 879ms | http |
| 120.92.212.16:8890 | ✓ 1181ms | ✓ 1466ms | ✓ 1217ms | ✓ 1514ms | ✓ 1201ms | http |
| 45.88.0.111:3128 | 否 | 否 | ✓ 1825ms | ✓ 1793ms | ✓ 1756ms | http |
| 45.88.0.115:3128 | 否 | 否 | ✓ 1823ms | ✓ 1819ms | ✓ 1460ms | http |
| 111.79.111.126:3128 | ✓ 1624ms | 否 | ✓ 1755ms | ✓ 1696ms | 否 | http |
| 45.88.0.113:3128 | ✓ 827ms | ✓ 1216ms | ✓ 889ms | 否 | 否 | http |
| 150.249.255.91:3128 | ✓ 1265ms | ✓ 1577ms | ✓ 781ms | ✓ 1408ms | ✓ 1854ms | http |
| 106.117.208.101:7890 | ✓ 1404ms | ✓ 1647ms | 否 | 否 | ✓ 1382ms | http |
| 38.180.2.107:3128 | ✓ 1332ms | ✓ 1793ms | ✓ 1928ms | 否 | ✓ 1918ms | http |
| 45.136.198.40:3128 | ✓ 1257ms | 否 | ✓ 940ms | 否 | ✓ 1470ms | http |
| 160.250.4.245:1 | ✓ 1902ms | 否 | 否 | ✓ 1654ms | ✓ 1232ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1539ms | 否 | ✓ 1589ms | ✓ 1535ms | http |
| 45.88.0.98:3128 | ✓ 1428ms | ✓ 1238ms | ✓ 1746ms | 否 | ✓ 1005ms | http |
| 45.136.130.245:8447 | ✓ 737ms | ✓ 1858ms | ✓ 921ms | ✓ 949ms | ✓ 721ms | http |
| 101.43.127.100:8877 | ✓ 1096ms | ✓ 1370ms | ✓ 1060ms | ✓ 1334ms | ✓ 1031ms | http |
| 45.140.147.82:1081 | ✓ 909ms | 否 | ✓ 958ms | ✓ 1774ms | ✓ 1633ms | http |
| 172.212.68.37:3128 | ✓ 392ms | 否 | ✓ 501ms | ✓ 1693ms | ✓ 1075ms | http |
| 198.24.188.140:22244 | 否 | 否 | ✓ 973ms | ✓ 1742ms | ✓ 1182ms | http |
| 103.113.70.189:1081 | ✓ 282ms | ✓ 1826ms | 否 | ✓ 1072ms | ✓ 807ms | http |
| 45.140.147.155:1081 | ✓ 1134ms | 否 | ✓ 970ms | ✓ 1608ms | ✓ 1409ms | http |

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
