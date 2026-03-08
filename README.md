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

最后更新：2026-03-08 12:15:32 UTC（2026-03-08 20:15:32 UTC+8）

**代理总数：46**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 45 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 46 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 1588ms | 否 | ✓ 1422ms | ✓ 1386ms | ✓ 1043ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1983ms | ✓ 1813ms | ✓ 1520ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 673ms | ✓ 1603ms | ✓ 636ms | http |
| 194.213.18.200:443 | ✓ 1629ms | 否 | ✓ 1060ms | 否 | ✓ 1805ms | http |
| 152.42.213.210:8080 | 否 | 否 | ✓ 1616ms | ✓ 1392ms | ✓ 1474ms | http |
| 120.232.242.119:22222 | ✓ 884ms | ✓ 1166ms | ✓ 1056ms | ✓ 1113ms | ✓ 900ms | http |
| 165.227.5.10:8888 | ✓ 668ms | ✓ 1040ms | ✓ 615ms | 否 | 否 | http |
| 47.77.193.180:1080 | ✓ 637ms | 否 | ✓ 472ms | ✓ 699ms | 否 | http |
| 120.92.212.16:8890 | 否 | ✓ 1425ms | 否 | ✓ 1215ms | ✓ 974ms | http |
| 202.155.12.161:443 | ✓ 1421ms | 否 | ✓ 1287ms | ✓ 1315ms | 否 | http |
| 81.70.169.194:80 | 否 | ✓ 1303ms | ✓ 1891ms | ✓ 1303ms | ✓ 1311ms | http |
| 115.231.181.40:8128 | ✓ 1526ms | ✓ 1193ms | 否 | 否 | ✓ 1978ms | http |
| 101.43.255.96:80 | 否 | ✓ 1271ms | 否 | ✓ 1520ms | ✓ 1323ms | http |
| 103.215.36.88:17170 | ✓ 946ms | ✓ 1826ms | ✓ 921ms | ✓ 1365ms | ✓ 1020ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1210ms | 否 | ✓ 1352ms | ✓ 1775ms | http |
| 121.230.8.49:1080 | ✓ 1388ms | ✓ 1586ms | 否 | ✓ 1472ms | ✓ 1158ms | http |
| 35.225.22.61:80 | ✓ 718ms | 否 | 否 | ✓ 1164ms | ✓ 1068ms | http |
| 46.249.103.192:443 | ✓ 730ms | 否 | ✓ 1420ms | 否 | ✓ 1628ms | http |
| 117.159.239.49:22222 | ✓ 776ms | ✓ 1011ms | ✓ 880ms | ✓ 1091ms | ✓ 832ms | http |
| 46.183.25.8:443 | ✓ 1637ms | 否 | ✓ 1211ms | 否 | ✓ 1598ms | http |
| 113.59.32.163:22222 | ✓ 1077ms | ✓ 1424ms | ✓ 886ms | 否 | 否 | http |
| 120.240.35.176:22222 | ✓ 959ms | ✓ 1392ms | ✓ 1092ms | 否 | 否 | http |
| 113.59.32.141:22222 | 否 | ✓ 1402ms | ✓ 1026ms | ✓ 1910ms | 否 | http |
| 211.171.114.154:3128 | ✓ 1956ms | 否 | ✓ 1747ms | ✓ 1781ms | ✓ 1826ms | http |
| 222.184.48.242:22222 | 否 | ✓ 1175ms | ✓ 1843ms | 否 | ✓ 876ms | http |
| 117.159.239.50:22222 | ✓ 809ms | ✓ 1027ms | ✓ 838ms | ✓ 1121ms | ✓ 843ms | http |
| 39.104.201.40:7890 | 否 | ✓ 1253ms | ✓ 1481ms | 否 | ✓ 1897ms | http |
| 210.223.44.230:3128 | ✓ 1510ms | ✓ 1247ms | 否 | 否 | ✓ 1756ms | http |
| 103.215.36.88:17254 | ✓ 1457ms | 否 | 否 | ✓ 1700ms | ✓ 1220ms | http |
| 85.9.195.140:1080 | ✓ 1198ms | 否 | ✓ 786ms | 否 | ✓ 1607ms | http |
| 150.249.255.91:3128 | ✓ 531ms | ✓ 1403ms | ✓ 562ms | ✓ 1559ms | 否 | http |
| 170.78.208.245:999 | ✓ 1621ms | 否 | ✓ 1797ms | 否 | ✓ 1686ms | http |
| 88.80.150.82:8080 | ✓ 1522ms | 否 | 否 | ✓ 1936ms | ✓ 1561ms | https |
| 183.249.5.111:22222 | ✓ 768ms | ✓ 1199ms | ✓ 763ms | ✓ 933ms | ✓ 724ms | http |
| 120.240.35.175:22222 | ✓ 871ms | ✓ 1205ms | ✓ 1161ms | 否 | 否 | http |
| 117.159.239.44:22222 | ✓ 1051ms | ✓ 1031ms | ✓ 860ms | 否 | ✓ 842ms | http |
| 120.198.141.79:22222 | ✓ 1231ms | ✓ 1502ms | ✓ 1030ms | ✓ 1229ms | ✓ 1086ms | http |
| 103.215.36.88:18516 | ✓ 1677ms | ✓ 1493ms | ✓ 1969ms | ✓ 1450ms | ✓ 1131ms | http |
| 120.240.35.177:22222 | ✓ 1017ms | ✓ 1283ms | ✓ 993ms | 否 | 否 | http |
| 172.212.68.37:3128 | ✓ 525ms | 否 | 否 | ✓ 1418ms | ✓ 1223ms | http |
| 183.249.5.214:22222 | 否 | ✓ 1099ms | ✓ 875ms | ✓ 1053ms | ✓ 905ms | http |
| 222.184.48.252:22222 | ✓ 1736ms | 否 | ✓ 1176ms | 否 | ✓ 1962ms | http |
| 103.39.51.190:8080 | ✓ 1968ms | 否 | ✓ 1801ms | ✓ 1868ms | 否 | http |
| 62.113.119.14:8080 | ✓ 763ms | ✓ 1735ms | ✓ 764ms | 否 | ✓ 1433ms | http |
| 45.136.198.40:3128 | ✓ 1291ms | ✓ 1979ms | ✓ 1308ms | 否 | 否 | http |
| 8.217.147.173:8080 | ✓ 1637ms | ✓ 1062ms | ✓ 980ms | ✓ 1124ms | ✓ 975ms | http |

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
