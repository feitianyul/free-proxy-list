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

最后更新：2026-03-14 18:23:47 UTC（2026-03-15 02:23:47 UTC+8）

**代理总数：43**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 42 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 43 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | ✓ 1568ms | ✓ 1964ms | ✓ 1324ms | ✓ 1507ms | ✓ 1185ms | http |
| 45.167.124.52:8080 | ✓ 1089ms | ✓ 1585ms | ✓ 1404ms | ✓ 1572ms | ✓ 1272ms | http |
| 205.209.118.30:3138 | ✓ 318ms | ✓ 1388ms | ✓ 1447ms | ✓ 1300ms | ✓ 791ms | http |
| 192.71.213.85:9812 | ✓ 1492ms | 否 | ✓ 1522ms | ✓ 1800ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1280ms | ✓ 1680ms | ✓ 1273ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 825ms | 否 | ✓ 1146ms | ✓ 1187ms | 否 | http |
| 38.145.203.135:8443 | ✓ 826ms | ✓ 1155ms | ✓ 580ms | ✓ 1379ms | ✓ 838ms | http |
| 38.145.218.82:8443 | ✓ 955ms | 否 | ✓ 281ms | ✓ 951ms | ✓ 866ms | http |
| 165.227.5.10:8888 | ✓ 1896ms | 否 | ✓ 1503ms | ✓ 1277ms | ✓ 1441ms | http |
| 101.43.255.96:80 | ✓ 1201ms | 否 | 否 | ✓ 1383ms | ✓ 1109ms | http |
| 167.71.196.28:8080 | ✓ 898ms | 否 | ✓ 1669ms | 否 | ✓ 1136ms | http |
| 45.149.92.147:5001 | ✓ 882ms | 否 | ✓ 858ms | ✓ 1045ms | ✓ 904ms | http |
| 85.198.96.242:3128 | ✓ 509ms | 否 | ✓ 910ms | ✓ 1624ms | ✓ 1490ms | http |
| 150.230.249.50:1080 | ✓ 1262ms | 否 | ✓ 1244ms | ✓ 1216ms | ✓ 1019ms | http |
| 103.82.23.118:5234 | ✓ 1957ms | 否 | ✓ 1495ms | 否 | ✓ 1877ms | http |
| 45.136.131.30:8447 | ✓ 659ms | ✓ 1305ms | ✓ 1180ms | ✓ 1047ms | ✓ 835ms | http |
| 45.136.131.31:8443 | 否 | 否 | ✓ 1118ms | ✓ 945ms | ✓ 679ms | http |
| 45.136.131.35:8443 | 否 | 否 | ✓ 1193ms | ✓ 936ms | ✓ 734ms | http |
| 62.60.177.204:34094 | 否 | ✓ 1213ms | ✓ 1082ms | 否 | ✓ 717ms | http |
| 150.249.255.91:3128 | ✓ 1645ms | ✓ 1179ms | ✓ 825ms | ✓ 1715ms | 否 | http |
| 107.174.55.123:7878 | ✓ 1529ms | ✓ 1458ms | ✓ 1148ms | ✓ 1121ms | ✓ 858ms | http |
| 194.5.212.40:8080 | ✓ 754ms | ✓ 1838ms | ✓ 1484ms | ✓ 1469ms | ✓ 1633ms | http |
| 138.124.53.221:443 | ✓ 773ms | 否 | ✓ 1896ms | ✓ 1759ms | 否 | http |
| 104.248.81.109:3128 | ✓ 468ms | 否 | 否 | ✓ 1644ms | ✓ 1163ms | http |
| 88.80.150.82:8080 | ✓ 1144ms | ✓ 1776ms | 否 | 否 | ✓ 1614ms | https |
| 210.223.44.230:3128 | 否 | ✓ 1279ms | ✓ 1212ms | 否 | ✓ 886ms | http |
| 81.70.169.194:80 | ✓ 1415ms | 否 | 否 | ✓ 1447ms | ✓ 1748ms | http |
| 34.101.184.164:3128 | ✓ 1206ms | 否 | ✓ 1358ms | ✓ 1768ms | ✓ 1174ms | http |
| 43.167.227.161:1080 | ✓ 679ms | ✓ 1551ms | 否 | ✓ 1073ms | ✓ 1141ms | http |
| 198.24.188.139:30143 | ✓ 1954ms | 否 | ✓ 1461ms | ✓ 1873ms | ✓ 1412ms | http |
| 14.225.212.37:7890 | ✓ 1948ms | 否 | ✓ 1820ms | ✓ 1327ms | 否 | http |
| 103.84.95.54:7890 | ✓ 1961ms | 否 | ✓ 979ms | ✓ 1217ms | ✓ 852ms | http |
| 106.117.208.101:7890 | ✓ 1174ms | 否 | ✓ 1299ms | 否 | ✓ 1971ms | http |
| 207.254.71.62:8088 | ✓ 1336ms | 否 | ✓ 1470ms | 否 | ✓ 1640ms | http |
| 210.77.29.245:7890 | ✓ 1049ms | 否 | ✓ 1373ms | ✓ 1657ms | ✓ 1374ms | http |
| 91.233.223.147:3128 | ✓ 1148ms | 否 | ✓ 1869ms | ✓ 1934ms | 否 | http |
| 45.129.141.143:3128 | ✓ 1360ms | 否 | ✓ 1965ms | 否 | ✓ 1709ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1326ms | 否 | ✓ 1293ms | ✓ 1557ms | http |
| 38.145.208.138:8447 | ✓ 366ms | ✓ 958ms | ✓ 1508ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 1076ms | ✓ 1825ms | ✓ 1984ms | ✓ 1744ms | ✓ 1423ms | http |
| 61.52.131.172:8443 | ✓ 1067ms | ✓ 1320ms | ✓ 1121ms | ✓ 1386ms | ✓ 1067ms | http |
| 38.145.208.246:8443 | ✓ 417ms | ✓ 963ms | ✓ 312ms | ✓ 1237ms | ✓ 732ms | http |
| 38.145.208.245:8443 | ✓ 390ms | ✓ 909ms | ✓ 688ms | ✓ 891ms | ✓ 700ms | http |

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
